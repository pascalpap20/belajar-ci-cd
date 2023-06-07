package repository

import (
	"crud/entity"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(dbCrud *gorm.DB) User {
	return User{
		db: dbCrud,
	}

}

type UserInterfaceRepo interface {
	CreateUser(user *entity.User) (*entity.User, error)
	GetUserById(id uint) (entity.User, error)
	UpdateUser(user entity.User) (any, error)
	DeleteUser(email string) (any, error)
}

// CreateUser new user
func (repo User) CreateUser(user *entity.User) (*entity.User, error) {
	err := repo.db.Model(&entity.User{}).Create(user).Error
	return user, err
}

// GetUserById get single user by id
func (repo User) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	repo.db.First(&user, "id = ? ", id)
	return user, nil
}

// UpdateUser multiple fields
func (repo User) UpdateUser(user entity.User) (any, error) {
	err := repo.db.Model(entity.User{}).
		Save(user).Error
	return nil, err
}

// DeleteUser by Id and email
func (repo User) DeleteUser(email string) (any, error) {
	err := repo.db.Model(&entity.User{}).
		Where("email = ?", email).
		Delete(&entity.User{}).
		Error
	return nil, err
}
