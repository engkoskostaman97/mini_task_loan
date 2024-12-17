package dto

type CreateConsumerRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	FullName     string  `json:"full_name" validate:"required"`
	LegalName    string  `json:"legal_name" validate:"required"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
	KTPPhoto     string  `json:"ktp_photo" validate:"required"`
	SelfiePhoto  string  `json:"selfie_photo" validate:"required"`
}

type UpdateConsumerRequest struct {
	NIK          string  `json:"nik" validate:"required"`
	FullName     string  `json:"full_name" validate:"required"`
	LegalName    string  `json:"legal_name" validate:"required"`
	PlaceOfBirth string  `json:"place_of_birth" validate:"required"`
	DateOfBirth  string  `json:"date_of_birth" validate:"required"`
	Salary       float64 `json:"salary" validate:"required"`
	KTPPhoto     string  `json:"ktp_photo" validate:"required"`
	SelfiePhoto  string  `json:"selfie_photo" validate:"required"`
}

type CreateLimitRequest struct {
	ConsumerID int     `json:"consumer_id"` 
	Tenor  int     `json:"tenor" validate:"required"`
	Amount float64 `json:"amount" validate:"required"`
}

type CreateTransactionRequest struct {
	ConsumerID int     `json:"consumer_id"` 
	ContractNo  string  `json:"contract_no" validate:"required"`
	OTR         float64 `json:"otr" validate:"required"`
	AdminFee    float64 `json:"admin_fee" validate:"required"`
	Installment float64 `json:"installment" validate:"required"`
	Interest    float64 `json:"interest" validate:"required"`
	AssetName   string  `json:"asset_name" validate:"required"`
}
