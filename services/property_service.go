package services

import (
	"refine-portal/models"
	"refine-portal/requests"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// GetProperties retrieves a list of properties.
//
// Responsibilities:
//   - Call the Property List request layer.
//   - Return the property list response.
//   - Log useful information for debugging.
func GetProperties(
    req models.PropertyListRequest,
) (*models.PropertyListResponse, error) {

	start := time.Now()

	defer func() {
		logs.Info(
			"[PropertyService] completed in %v",
			time.Since(start),
		)
	}()

    logs.Debug(
        "[PropertyService] Fetching property list | category=%s | location=%s",
        req.Category,
        req.Locations,
    )

    response, err := requests.GetPropertyListRequest(req)
    if err != nil {
        return nil, err
    }

    logs.Info(
        "[PropertyService] Total Properties=%d",
        len(response.Result.ItemIDs),
    )

    return response, nil
}
