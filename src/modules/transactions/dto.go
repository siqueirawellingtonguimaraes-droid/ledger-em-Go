package transactions

type TransactionEntryCreateDTO struct {
	AccountID string `json:"account_id" binding:"required"`
	Amount    int64  `json:"amount" binding:"required"`
	Currency  string `json:"currency" binding:"required,len=3"`
	Direction string `json:"direction" binding:"required"`
}

type TransactionCreateDTO struct {
	Description string                      `json:"description"`
	Entries     []TransactionEntryCreateDTO `json:"entries" binding:"required,min=2"`
}

type TransactionEntryResponseDTO struct {
	ID        string `json:"id"`
	AccountID string `json:"account_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
	Direction string `json:"direction"`
}

type TransactionResponseDTO struct {
	ID          string                        `json:"id"`
	Description string                        `json:"description"`
	Entries     []TransactionEntryResponseDTO `json:"entries"`
}
