package models

import "time"

// Consumer 
type Consumer struct {
	ID           uint      `gorm:"primaryKey"`
	NIK          string    `gorm:"unique;not null"`
	FullName     string    `gorm:"not null"`
	LegalName    string    `gorm:"not null"`
	PlaceOfBirth string    `gorm:"not null"`
	DateOfBirth  time.Time `gorm:"not null"`
	Salary       float64   `gorm:"not null"`
	KTPPhoto     string    `gorm:"not null"`
	SelfiePhoto  string    `gorm:"not null"`
	Limits       []Limit   `gorm:"foreignKey:ConsumerID"`
	Transactions []Transaction `gorm:"foreignKey:ConsumerID"`
}

// Limit 
type Limit struct {
	ID         uint    `gorm:"primaryKey"`
	ConsumerID uint    `gorm:"not null"`
	Tenor      int     `gorm:"not null"`  
	Amount     float64 `gorm:"not null"`  
}

// Transaction 
type Transaction struct {
	ID           uint      `gorm:"primaryKey"`
	ConsumerID   uint      `gorm:"not null"`
	ContractNo   string    `gorm:"unique;not null"`
	OTR          float64   `gorm:"not null"`
	AdminFee     float64   `gorm:"not null"`
	Installment  float64   `gorm:"not null"`
	Interest     float64   `gorm:"not null"`
	AssetName    string    `gorm:"not null"`
}
