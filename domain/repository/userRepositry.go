package repository

import (
	"kunikida123456/NutritionApp/domain/model"
)

type UserRepository interface {
	CreateUser(meal *model.User) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
}
