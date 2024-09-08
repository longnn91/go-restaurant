package usecases

import (
	"gogo/internal/app/entities"
	"gogo/internal/repository"
	"gogo/internal/utils"
)

type UserUsecase struct {
	Repo repository.UserRepository
}

func (u *UserUsecase) Register(username, password string) (*entities.User, error) {
	//Handle to save user with hashed password
	var hashPassword utils.Password
	if err := hashPassword.Set(password); err != nil {
		return nil, err
	}

	hashedPassword := hashPassword.Hash

	user := &entities.User{
		Username: username,
		Password: hashedPassword,
	}

	err := u.Repo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserUsecase) Login(username, password string) (string, error) {
	user, err := u.Repo.FindByUsername(username)
	if err != nil {
		return "", err
	}

	var hashPassword utils.Password
	if err := hashPassword.Set(password); err != nil {
		return "", err
	}

	if hashPassword.Matches(user.Password); err != nil {
		return "", err
	}

	token, err := utils.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
