package services

import (
	"refine-portal/models"
	"refine-portal/requests"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// GetCategory retrieves category page data.
//
// Responsibilities:
//   - Call the Category request layer.
//   - Return the category response.
//   - Log execution time.
func GetCategory(
	slug string,
	countryCode string,
) (*models.CategoryResponse, error) {

	start := time.Now()

	defer func() {
		logs.Info(
			"[CategoryService] completed in %v",
			time.Since(start),
		)
	}()

	return requests.GetCategoryRequest(
		slug,
		countryCode,
	)
}