package models

type PropertyImagesResponse struct {
	Success bool     `json:"Success"`
	Message string   `json:"Message"`
	Error   any      `json:"Error"`
	Images  []string `json:"Images"`
}