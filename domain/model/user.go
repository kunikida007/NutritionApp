package model

import (
	"errors"
	"time"
)

type User struct {
	ID        int       `json:"user_id" gorm:"praimaly_key"`
	Meals     []Meal    `json:"meals" gorm:"foreignKey:UserID"`
	Name      string    `json:"name"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"password" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}

func NewUser(name string, email string, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	if email == "" {
		return nil, errors.New("email is empty")
	}
	user := &User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	return user, nil
}

// func (u *User) Set(name string, email string, password string) error {
// 	if name == "" {
// 		return errors.New("name is empty")
// 	}
// 	if email == "" {
// 		return errors.New("email is empty")
// 	}

// 	u.Name = name
// 	u.Email = email
// 	u.Password = password

// 	return nil
// }
