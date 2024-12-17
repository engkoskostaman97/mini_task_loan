package repository

import (
	"xyz_multifinance/database"
	"xyz_multifinance/models"
)

func CreateTransaction(transaction *models.Transaction) error {
	return database.DB.Create(transaction).Error
}

func GetAllTransactions(transactions *[]models.Transaction) error {
	return database.DB.Find(transactions).Error
}

func GetTransactionByID(id uint, transaction *models.Transaction) error {
	return database.DB.First(transaction, id).Error
}

func UpdateTransaction(transaction *models.Transaction) error {
	return database.DB.Save(transaction).Error
}

func DeleteTransaction(id uint) error {
	return database.DB.Delete(&models.Transaction{}, id).Error
}
