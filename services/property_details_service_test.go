package services

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"

	"refine-portal/models"
	"refine-portal/requests"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestGetPropertyDetails_Success(t *testing.T) {

	req := models.PropertyDetailsRequest{
		PropertyIDList: []string{"101", "102"},
	}

	expectedBatch := &models.PropertyDetailsResponse{
		Success: true,

		Items: []models.PropertyDetails{
			{
				ID: "101",
				Property: models.Property{
					FeatureImage: "house1.jpg",
				},
			},
			{
				ID: "102",
				Property: models.Property{
					FeatureImage: "house2.jpg",
				},
			},
		},

		Result: models.PropertyDetailsResult{
			ItemsByID: map[string]models.PartnerInfo{
				"101": {
					Feed: 11,
				},
				"102": {
					Feed: 12,
				},
			},
		},
	}

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetPropertyDetailsRequest,
		func(ids []string) (*models.PropertyDetailsResponse, error) {

			assert.Equal(t, req.PropertyIDList, ids)

			return expectedBatch, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(key string) (string, error) {

			assert.Equal(t, "image_base_url", key)

			return "https://images.test.com", nil
		},
	)

	defer patches.Reset()

	result, err := GetPropertyDetails(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Len(t, result.Items, 2)

	assert.Equal(
		t,
		"https://images.test.com/house1.jpg",
		result.Items[0].Property.FeatureImage,
	)

	assert.Equal(
		t,
		"https://images.test.com/house2.jpg",
		result.Items[1].Property.FeatureImage,
	)

	assert.Equal(
		t,
		11,
		result.Items[0].Feed,
	)

	assert.Equal(
		t,
		12,
		result.Items[1].Feed,
	)
}


func TestGetPropertyDetails_RequestError(t *testing.T) {

	req := models.PropertyDetailsRequest{
		PropertyIDList: []string{"101", "102"},
	}

	expectedErr := errors.New("request failed")

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetPropertyDetailsRequest,
		func(ids []string) (*models.PropertyDetailsResponse, error) {
			return nil, expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetPropertyDetails(req)

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestGetPropertyDetails_ConfigError(t *testing.T) {

	req := models.PropertyDetailsRequest{
		PropertyIDList: []string{"101"},
	}

	expectedBatch := &models.PropertyDetailsResponse{
		Success: true,
		Items: []models.PropertyDetails{
			{
				ID: "101",
				Property: models.Property{
					FeatureImage: "house.jpg",
				},
			},
		},
		Result: models.PropertyDetailsResult{
			ItemsByID: map[string]models.PartnerInfo{
				"101": {
					Feed: 11,
				},
			},
		},
	}

	expectedErr := errors.New("config error")

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetPropertyDetailsRequest,
		func(ids []string) (*models.PropertyDetailsResponse, error) {
			return expectedBatch, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(key string) (string, error) {
			return "", expectedErr
		},
	)

	defer patches.Reset()

	result, err := GetPropertyDetails(req)

	assert.Nil(t, result)
	assert.Equal(t, expectedErr, err)
}

// Verifies that property IDs are split into batches and
// each batch is processed concurrently before results are merged.
func TestGetPropertyDetails_ConcurrentBatchProcessing(t *testing.T) {

	req := models.PropertyDetailsRequest{}

	for i := 1; i <= 55; i++ {
		req.PropertyIDList = append(
			req.PropertyIDList,
			fmt.Sprintf("%d", i),
		)
	}

	var (
		batchSizes []int
		mu         sync.Mutex
	)
	var callCount atomic.Int32

	patches := gomonkey.NewPatches()

	patches.ApplyFunc(
		requests.GetPropertyDetailsRequest,
		func(ids []string) (*models.PropertyDetailsResponse, error) {

			callCount.Add(1)

			mu.Lock()
			batchSizes = append(batchSizes, len(ids))
			mu.Unlock()

			return &models.PropertyDetailsResponse{
				Success: true,
			}, nil
		},
	)

	patches.ApplyFunc(
		requests.GetURLFromConfig,
		func(key string) (string, error) {
			return "https://images.test.com", nil
		},
	)

	defer patches.Reset()

	result, err := GetPropertyDetails(req)

	assert.NoError(t, err)
	assert.NotNil(t, result)

	assert.Equal(t, int32(2), callCount.Load())
	assert.Len(t, batchSizes, 2)
	assert.Contains(t, batchSizes, 50)
	assert.Contains(t, batchSizes, 5)
}