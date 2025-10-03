package repository

// import (
// 	"gorm.io/gorm"
// )

// type Repository struct {
// 	User         UserRepository
// 	Event        EventRepository
// 	EventPicture eventPictureRepository
// }

// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{
// 		User:         NewUserRepository(db),
// 		Event:        NewEventRepository(db),
// 		EventPicture: NewEventPictureRepository(db),
// 	}
// }