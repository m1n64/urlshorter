package services

type AddRequest struct {
	Link string `json:"link" form:"link"`
}

type AddResponse struct {
	Status bool   `json:"status"`
	Slug   string `json:"slug"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Status bool   `json:"status"`
	Token  string `json:"token"`
}
