package services

import (
	"refine-portal/models"
	"refine-portal/requests"
	"sync"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

const (
	propertyDetailsAPIPath = "/api/property/bookmark/v1"
	batchSize              = 50
)

// GetPropertyDetails retrieves property details for a list of property IDs.
//
// Responsibilities:
//   - Split property IDs into batches.
//   - Fetch all batches concurrently.
//   - Preserve the original batch order.
//   - Merge all batch responses into a single result.
// 	 - Build image URLs.
//   - Return the combined property details response.
func GetPropertyDetails(
	req models.PropertyDetailsRequest,
) (*models.PropertyDetailsResponse, error) {

	start := time.Now()

	defer func() {
		logs.Info(
			"[PropertyDetailsService] completed in %v",
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

	batchResults := make(
		chan batchResult,
		len(chunks),
	)

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

			batch, err := requests.GetPropertyDetailsRequest(propertyIDs)

			batchResults <- batchResult{
				Index: idx,
				Data: batch,
				Err: err,
			}
		}(index, ids)
	}

	wg.Wait()
	close(batchResults)

	orderedBatches := make(
		[]*models.PropertyDetailsResponse,
		len(chunks),
	)

	for result := range batchResults {
		if result.Err != nil {
			return nil, result.Err
		}

		orderedBatches[result.Index] = result.Data
	}

	// Merge batch responses while preserving the original order.
	for _, batch := range orderedBatches {

		if batch == nil {
			continue
		}

		merged.Items = append(
			merged.Items,
			batch.Items...,
		)

		for id, info := range batch.Result.ItemsByID {
			merged.Result.ItemsByID[id] = info
		}
	}

	imageBaseURL, err := requests.GetURLFromConfig("image_base_url")
	if err != nil {
		return nil, err
	}

	// Enrich properties with full image URLs and partner feed information.
	for i := range merged.Items {

		image := merged.Items[i].Property.FeatureImage

		if image != "" {
			merged.Items[i].Property.FeatureImage =
				requests.BuildImageURL(
					imageBaseURL,
					image,
				)
		}

		// Preserve the original property ID order when attaching partner info.
		if i < len(req.PropertyIDList) {

			propertyID := req.PropertyIDList[i]

			if partnerInfo, ok := merged.Result.ItemsByID[propertyID]; ok {
				merged.Items[i].Feed = partnerInfo.Feed
			}
		}
	}

	logs.Info(
		"[PropertyDetailsService] Total Properties=%d",
		len(merged.Items),
	)

	return merged, nil
}

