package services

import (
	"refine-portal/models"
	"refine-portal/requests"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// GetLocation retrieves matching locations.
//
// Responsibilities:
//   - Call the Location request layer.
//   - Return the location response.
//   - Log total execution time.
func GetLocation(
    keyword string,
) (*models.LocationResponse, error) {

    start := time.Now()

    defer func() {
        logs.Info(
            "[LocationService] completed in %v",
            time.Since(start),
        )
    }()

    return requests.GetLocationRequest(keyword)
}
