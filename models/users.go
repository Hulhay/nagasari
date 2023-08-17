package models

type User struct {
	ID          int64   `db:"id"`
	UUID        string  `db:"uuid"`
	Name        string  `db:"name"`
	Email       string  `db:"email"`
	Password    string  `db:"password"`
	PhoneNumber string  `db:"phone_number"`
	Address     *string `db:"address"`
	PhotoURL    string  `db:"photo_url"`
	CreatedAt   string  `db:"created_at"`
}
