package services

import (
	"refine-portal/models"
	"refine-portal/requests"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
)

// GetCategory retrieves category page data.
//
// Responsibilities:
//   - Call the Category request layer.
//   - Replace location placeholders in section titles.
//   - Convert feature image filenames into full image URLs.
//   - Return the processed category response.
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

	category, err := requests.GetCategoryRequest(
		slug,
		countryCode,
	)
	if err != nil {
		return nil, err
	}

	// Replace "{{.Location}}" placeholders in section titles
	// with the current location name returned by the API.
	locationName := category.GeoInfo.ShortName

	for i := range category.Result.Sections {
		category.Result.Sections[i].Title = 
		strings.ReplaceAll(
			category.Result.Sections[i].Title,
			"{{.Location}}",
			locationName,
		)

		category.Result.Sections[i].SubTitle = 
		strings.ReplaceAll(
			category.Result.Sections[i].SubTitle,
			"{{.Location}}",
			locationName,
		)
	}

	// Build full image URLs for all feature images.
	imageBaseURL, err := requests.GetURLFromConfig("image_base_url")
	if err != nil {
		return nil, err
	}

	for i := range category.Result.Sections {

		for j := range category.Result.Sections[i].Items {

			image :=
				category.Result.Sections[i].Items[j].Property.FeatureImage

			if image == "" {
				continue
			}

			category.Result.Sections[i].Items[j].Property.FeatureImage =
				requests.BuildImageURL(
					imageBaseURL,
					image,
				)
		}
	}

	return category, nil
}