package dto

type CreateProductInput struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenUserJwtInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GenUserJwtOutput struct {
	AccessToken string `json:"access_token"`
}
