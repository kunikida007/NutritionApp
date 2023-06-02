package infra

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"

	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserRepository struct {
	Conn *gorm.DB
}

func NewUserRepository(conn *gorm.DB) repository.UserRepository {
	return &UserRepository{Conn: conn}
}

func (ur *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	if err := ur.Conn.Create(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	u := model.User{}
	err := ur.Conn.Where("email = ?", email).First(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return &model.User{}, nil
		}
		return nil, err
	}
	return &u, nil
}

// func (ur *UserRepository) UserExists(uid int) error {
// 	u := model.User{}
// 	if err := ur.Conn.First(&u, uid).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
