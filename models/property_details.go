package models

type PropertyDetailsRequest struct {
	PropertyIDList []string
}

type PropertyDetailsResponse map[string]any