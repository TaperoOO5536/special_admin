package repository

// import (
// 	"gorm.io/gorm"
// )

// type Repository struct {
// 	User         UserRepository
// 	Ivent        IventRepository
// 	IventPicture iventPictureRepository
// }

// func NewRepository(db *gorm.DB) *Repository {
// 	return &Repository{
// 		User:         NewUserRepository(db),
// 		Ivent:        NewIventRepository(db),
// 		IventPicture: NewIventPictureRepository(db),
// 	}
// }