package user

import (
	"crud/entity"
	"crud/repository"
	"time"
)

type UseCaseUser interface {
	CreateUser(user UserParam) (entity.User, error)
	GetUserById(id uint) (entity.User, error)
	UpdateUser(user UserParam, id uint) (any, error)
	DeleteUser(email string) (any, error)
}

type useCaseUser struct {
	userRepo repository.UserInterfaceRepo
}

func (uc useCaseUser) CreateUser(user UserParam) (entity.User, error) {
	var newUser *entity.User

	newUser = &entity.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := uc.userRepo.CreateUser(newUser)
	if err != nil {
		return *newUser, err
	}
	return *newUser, nil
}

func (uc useCaseUser) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	user, err := uc.userRepo.GetUserById(id)
	return user, err
}

func (uc useCaseUser) UpdateUser(user UserParam, id uint) (any, error) {
	var editUser entity.User
	editUser = entity.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		UpdatedAt: time.Now(),
	}

	_, err := uc.userRepo.UpdateUser(editUser)
	if err != nil {
		return editUser, err
	}
	return editUser, nil
}

func (uc useCaseUser) DeleteUser(email string) (any, error) {
	_, err := uc.userRepo.DeleteUser(email)
	return nil, err
}
