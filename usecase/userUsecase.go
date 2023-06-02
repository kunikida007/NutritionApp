package usecase

import (
	"kunikida123456/NutritionApp/domain/model"
	"kunikida123456/NutritionApp/domain/repository"
	"kunikida123456/NutritionApp/myerror"
	"kunikida123456/NutritionApp/util"
)

type UserUsecase interface {
	Signup(name, email, password string) (*model.User, error)
	Login(email, password string) (string, *model.User, error)
}

type userUsecase struct {
	repository repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repository: userRepo,
	}
}

func (uu *userUsecase) Signup(name, email, password string) (*model.User, error) {
	exsitUser, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}
	if exsitUser.ID != 0 {
		return nil, &myerror.BadRequestError{Msg: "user already exist"}
	}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	user, err := model.NewUser(name, email, hashedPassword)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	createdUser, err := uu.repository.CreateUser(user)
	if err != nil {
		return nil, &myerror.BadRequestError{Err: err}
	}

	return createdUser, nil
}

func (uu *userUsecase) Login(email, password string) (string, *model.User, error) {
	user, err := uu.repository.GetUserByEmail(email)
	if err != nil {
		return "", nil, &myerror.BadRequestError{Err: err}
	}
	if user.ID == 0 {
		return "", nil, &myerror.NotFoundError{Msg: "user not found or password is invalid"}
	}

	err = util.CheckPassword(user.Password, password)
	if err != nil {
		return "", nil, &myerror.NotFoundError{Err: err, Msg: "user not found or password is invalid"}
	}

	signedString, err := util.GenerateSignedString(user.ID, user.Name)
	if err != nil {
		return "", nil, &myerror.BadRequestError{Err: err}
	}

	return signedString, user, nil
}
