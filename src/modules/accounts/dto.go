package accounts

type AccountCreateDTO struct {
	Name string `json:"name" validate:"required"`
}

type AccountResponseDTO struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
