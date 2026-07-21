package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"refine-portal/models"
	"strings"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

const (
	propertyDetailsAPIPath = "/api/property/bookmark/v1"
	batchSize              = 50
)

func GetPropertyDetails(
	req models.PropertyDetailsRequest,
) (*models.PropertyDetailsResponse, error) {

	start := time.Now()

	defer func() {
		logs.Info(
			"[PropertyDetailsService] total took %v",
			time.Since(start),
		)
	}()

	chunks := chunkStrings(
		req.PropertyIDList,
		batchSize,
	)

	logs.Debug(
		"[PropertyDetailsService] Total IDs=%d | Total Chunks=%d",
		len(req.PropertyIDList),
		len(chunks),
	)

	merged := &models.PropertyDetailsResponse{
		Success: true,
		Result: models.PropertyDetailsResult{
			ItemsByID: make(map[string]models.PartnerInfo),
		},
	}

	type batchResult struct {
		Index int
		Data *models.PropertyDetailsResponse
		Err error
	}

	results := make(chan batchResult, len(chunks))

	var wg sync.WaitGroup

	for index, ids := range chunks {

		wg.Add(1)

		go func(idx int, propertyIDs []string) {
			defer wg.Done()

			logs.Debug(
				"[PropertyDetailsService] Processing batch %d/%d",
				idx+1,
				len(chunks),
			)

			batch, err := fetchPropertyDetailsBatch(propertyIDs)

			results <- batchResult{
				Index: idx,
				Data: batch,
				Err: err,
			}
		} (index, ids)
	}

	wg.Wait()
	close(results)

	batches := make([]*models.PropertyDetailsResponse, len(chunks))

	for result := range results {
		if result.Err != nil {
			return nil, result.Err
		}

		batches[result.Index] = result.Data
	}

	// Merge in original order
	for _, batch := range batches {
		merged.Items = append(merged.Items, batch.Items...)

		for id, info := range batch.Result.ItemsByID {
			merged.Result.ItemsByID[id] = info
		}
	}

	logs.Info(
		"[PropertyDetailsService] Total Properties=%d",
		len(merged.Items),
	)

	return merged, nil
}

func fetchPropertyDetailsBatch(
	propertyIDs []string,
) (*models.PropertyDetailsResponse, error) {

	start := time.Now()

	defer func() {
		logs.Info(
			"[PropertyDetailsBatch] size=%d took %v",
			len(propertyIDs),
			time.Since(start),
		)
	}()

	baseURL, err := web.AppConfig.String("base_url")
	if err != nil {
		logs.Error(
			"[PropertyDetailsService] Failed to read configuration | key=base_url | err=%v",
			err,
		)
		return nil, fmt.Errorf("failed to get 'base_url' from config: %w", err)
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("parse base_url failed: %w", err)
	}

	parsedURL.Path = propertyDetailsAPIPath

	query := parsedURL.Query()

	query.Set(
		"propertyIdList",
		strings.Join(propertyIDs, ","),
	)

	parsedURL.RawQuery = query.Encode()

	logs.Debug(
		"[PropertyDetailsService] Calling Property Details API | propertyIdCount=%d | url=%s",
		len(propertyIDs),
		parsedURL.String(),
	)

	request, err := NewGETRequest(
		parsedURL.String(),
	)
	if err != nil {
		return nil, err
	}

	response, err := httpClient.Do(request)
	if err != nil {
		logs.Error(
			"[PropertyDetailsService] HTTP request failed | url=%s | err=%v",
			parsedURL.String(),
			err,
		)
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {

		logs.Warn(
			"[PropertyDetailsService] Unexpected response | status=%d | url=%s",
			response.StatusCode,
			parsedURL.String(),
		)

		return nil, fmt.Errorf(
			"unexpected status code: %d",
			response.StatusCode,
		)

	}

	var result models.PropertyDetailsResponse

	if err := json.NewDecoder(
		response.Body,
	).Decode(&result); err != nil {
		logs.Error(
			"[PropertyDetailsService] Decode response failed | err=%v",
			err,
		)
		return nil, fmt.Errorf("decode response failed: %w", err)
	}

	return &result, nil
}
