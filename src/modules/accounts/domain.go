package accounts

type Account struct {
	ID   string
	Name string
}

type AccountWithBalance struct {
	Account Account
	Balance float64
}
