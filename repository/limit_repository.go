package repository

import (
	"xyz_multifinance/database"
	"xyz_multifinance/models"
)

func CreateLimit(limit *models.Limit) error {
	return database.DB.Create(limit).Error
}
