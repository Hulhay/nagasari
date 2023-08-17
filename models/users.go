package models

type User struct {
	ID          int64   `db:"id"`
	UUID        string  `db:"uuid"`
	Name        string  `db:"name"`
	Email       string  `db:"email"`
	PhoneNumber string  `db:"phone_number"`
	Address     *string `db:"address"`
	PhotoURL    string  `db:"photo_url"`
	CreatedAt   string  `db:"created_at"`
}

type RegisterRequest struct {
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	Password          string `json:"password"`
	EncryptedPassword string
}
