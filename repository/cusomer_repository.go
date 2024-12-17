package repository

import (
	"xyz_multifinance/database"
	"xyz_multifinance/models"
)

func CreateConsumer(consumer *models.Consumer) error {
	return database.DB.Create(consumer).Error
}

func GetAllConsumers(consumers *[]models.Consumer) error {
	return database.DB.Find(consumers).Error
}

func GetConsumerByID(id uint, consumer *models.Consumer) error {
	return database.DB.First(consumer, id).Error
}

func UpdateConsumer(consumer *models.Consumer) error {
	return database.DB.Save(consumer).Error
}

func DeleteConsumer(id uint) error {
	return database.DB.Delete(&models.Consumer{}, id).Error
}
