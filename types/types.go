package types

import (
	"time"
)

// User struct with improved data types
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

// Password struct with improved data types
type Password struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Quantity    int       `json:"quantity"`
	CreatedAt   time.Time `json:"created_at"`
}

// Orders struct with improved data types
type Order struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
	Total     int       `json:"total"`
	Status    string    `json:"status"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
}

// Orders_Items struct with improved data types
type OrderItem struct {
	ID        string  `json:"id"`
	OrderID   string  `json:"order_id"`
	ProductID string  `json:"product_id"`
	Quantity  int     `json:"quantity"`
	Price     float32 `json:"price"`
}

// RegisterUserPayload struct for JSON request payload
type RegisterUserPayload struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	EmailID   string `json:"email-id"`
	Password  string `json:"password"`
}

// UserDB defines the interface for user database operations
type UserDB interface {
	CreateUser(user User) (string, error)
	GetUserByID(id string) (*User, error)
	GetUserByEmail(email string) (*User, error)
	UpdateUser(user User) error
	DeleteUser(id int64) error
}

// PasswordDB defines the interface for password database operations
type PasswordDB interface {
	CreatePassword(password Password) (int64, error)
	GetPasswordByID(id int64) (*Password, error)
	UpdatePassword(password Password) error
	DeletePassword(id int64) error
}

// OrderDB defines the interface for order database operations
type OrderDB interface {
	CreateOrder(order Order) (int64, error)
	GetOrderByID(id int64) (*Order, error)
	GetOrdersByUserID(userID int64) ([]Order, error)
	UpdateOrder(order Order) error
	DeleteOrder(id int64) error
}

// OrderItemDB defines the interface for order item database operations
type OrderItemDB interface {
	CreateOrderItem(orderItem OrderItem) (int64, error)
	GetOrderItemByID(id int64) (*OrderItem, error)
	GetOrderItemsByOrderID(orderID int64) ([]OrderItem, error)
	UpdateOrderItem(orderItem OrderItem) error
	DeleteOrderItem(id int64) error
}
