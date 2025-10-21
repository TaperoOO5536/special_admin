package models

import (
	"time"

	"github.com/google/uuid"
)

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
}

type PaginatedEvents struct {
	Events     []Event `json:"events"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type PaginatedOrders struct {
	Orders     []Order `json:"events"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type PaginatedItems struct {
	Items     []Item `json:"events"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type PaginatedUsers struct {
	Users      []User `json:"events"`
	TotalCount int64   `json:"total_count"`
	Page       int     `json:"page"`
	PerPage    int     `json:"per_page"`
}

type Admin struct {
	Login            string    `gorm:"column:admin_login;primaryKey"`
	PasswordHash     string	   `gorm:"column:admin_password_hash"`
	RefreshTokenHash string    `gorm:"column:refresh_token_hash"`
	RefreshExpiresAt time.Time `gorm:"column:refresh_expires_at"`
}

type User struct {
	ID          string      `gorm:"column:id_user;primaryKey"`
	Name        string      `gorm:"column:f_n_user"`
	Surname     string      `gorm:"column:s_n_user"`
	Nickname    string      `gorm:"column:n_n_user"`
	PhoneNumber string      `gorm:"column:phone_n_user"`
	UserEvents  []UserEvent `gorm:"foreignKey:UserID"`
	Orders      []Order     `gorm:"foreignKey:UserID"`
}

type Event struct {
	ID            uuid.UUID      `gorm:"column:id_event;primaryKey"`
	Title         string         `gorm:"column:event_title"`
	Description   string         `gorm:"column:event_description"`
	DateTime      time.Time      `gorm:"column:event_datetime"`
	Price         int64          `gorm:"column:event_price"`
	TotalSeats    int64          `gorm:"column:total_seats"`
	OccupiedSeats int64          `gorm:"column:occupied_seats"`
	LittlePicture []byte         `gorm:"column:little_picture"`
	MimeType			string         `gorm:"column:mime_type"`
	UserEvents    []UserEvent    `gorm:"foreignKey:EventID"`
	Pictures      []EventPicture `gorm:"foreignKey:EventID"`
}

type UserEvent struct {
	ID             uuid.UUID `gorm:"column:id_user_event;primaryKey"`
	UserID         string    `gorm:"column:user_id"`
	EventID        uuid.UUID `gorm:"column:event_id"`
	NumberOfGuests int64     `gorm:"column:number_of_guests"`
	User           User      `gorm:"foreignKey:UserID"`
}

type EventPicture struct {
	ID       uuid.UUID `gorm:"column:id_event_picture;primaryKey"`
	EventID  uuid.UUID `gorm:"column:event_id"`
	Path     []byte    `gorm:"column:picture_path"`
	MimeType string   `gorm:"column:mime_type"`
}

type Order struct {
	ID             uuid.UUID   `gorm:"column:id_order;primaryKey"`
	Number         string      `gorm:"column:order_number"`
	UserID         string      `gorm:"column:user_id"`
	FormDate       time.Time   `gorm:"column:order_form_datetime"`
	CompletionDate time.Time   `gorm:"column:completion_date"`
	Comment        string      `gorm:"column:order_comment"`
	Status         string      `gorm:"column:order_status"`
	OrderAmount    int64       `gorm:"column:order_amount"`
	OrderItems     []OrderItem `gorm:"foreignKey:OrderID"`
	User           User        `gorm:"foreignKey:UserID"`
}

type Item struct {
	ID            uuid.UUID     `gorm:"column:id_item;primaryKey"`
	Title         string        `gorm:"column:item_title"`
	Description   string        `gorm:"column:item_description"`
	Price         int64         `gorm:"column:item_price"`
	LittlePicture []byte        `gorm:"column:little_picture"`
	MimeType			string        `gorm:"column:mime_type"`
	Pictures      []ItemPicture `gorm:"foreignKey:ItemID"`
}

type OrderItem struct{
	ID       uuid.UUID `gorm:"column:id_order_item;primaryKey"`
	OrderID  uuid.UUID `gorm:"column:order_id;index"`
	ItemID   uuid.UUID `gorm:"column:item_id;index"`
	Quantity int64
	Item 	   Item      `gorm:"foreignKey:ItemID"`
}

type ItemPicture struct {
	ID       uuid.UUID `gorm:"column:id_item_picture;primaryKey"`
	ItemID   uuid.UUID `gorm:"column:item_id"`
	Path     []byte    `gorm:"column:picture_path"`
	MimeType string  `gorm:"column:mime_type"`
}