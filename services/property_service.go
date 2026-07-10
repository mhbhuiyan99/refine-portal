package services

import "refine-portal/models"


func GetProperties(category, countryCode string, order, page int) (*models.PropertyListResponse, error)