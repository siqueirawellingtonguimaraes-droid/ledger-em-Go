package vo

type Currency string
type Direction string

const (
	BRL Currency = "BRL"
	USD Currency = "USD"
	EUR Currency = "EUR"
)

const (
	Credit Direction = "credit"
	Debit  Direction = "debit"
)
