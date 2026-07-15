package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"refine-portal/models"

	beego "github.com/beego/beego/v2/server/web"
)

func GetCategory(slug string) (*models.CategoryResponse, error) {

	baseURL, _ := beego.AppConfig.String("base_url")

	url := fmt.Sprintf(
		"%s/api/v1/category/details/%s?aggsAvgPrice=1&aggsAvgRating=1&aggsAvgRoomSize=1&aggsCategory=1&device=desktop&items=1&locations=US&sections=1",
		baseURL,
		slug,
	)

	request, err := NewGETRequest(url)
	if err != nil {
		return nil, err
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(
			"category api returned %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	var result models.CategoryResponse

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}