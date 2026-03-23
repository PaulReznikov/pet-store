package models

type User struct {
	ID        int           `json:"id"`
	Name      string        `json:"name"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	Role      UserRole      `json:"role"`
	Phone     string        `json:"phone"`
	Address   []UserAddress `json:"address"`
	Cards     []UserCards   `json:"cards"`
	CreatedAt string        `json:"created_at"`
	UpdatedAt string        `json:"updated_at"`
}

type UserRole struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserAddress struct {
	ID     int    `json:"id"`
	UserID int    `json:"user_id"`
	Street string `json:"street"`
	City   string `json:"city"`
}

type UserCards struct {
	ID         int    `json:"id"`
	UserID     int    `json:"user_id"`
	CardNumber string `json:"card_number"`
	ExpiryDate string `json:"expiry_date"`
	CVV        string `json:"cvv"`
}
