package models

type RegisterRequest struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password"`
	EncryptedPassword string
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token     string `json:"token"`
	ExpiredAt string `json:"expired_at"`
}
