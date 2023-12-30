package repository

import (
	. "crm-app-go/model"

	"github.com/jinzhu/gorm"
)

type IFlywayRepository interface {
	GetUserById(id int) (*User, error)
	GetAllUsers() (*[]User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) error
}

type flywayRepository struct {
	DB *gorm.DB
}

func NewFlywayRepository(db *gorm.DB) IFlywayRepository {
	return &flywayRepository{db}
}

func (flywayRepository *flywayRepository) GetUserById(id int) (*User, error) {
	var user User
	result := flywayRepository.DB.First(&user, id)
	return &user, result.Error
}

func (flywayRepository *flywayRepository) GetAllUsers() (*[]User, error) {
	var user []User
	result := flywayRepository.DB.Find(&user)
	return &user, result.Error
}

func (flywayRepository *flywayRepository) CreateUser(user *User) (*User, error) {
	result := flywayRepository.DB.Create(user)
	return user, result.Error
}

func (flywayRepository *flywayRepository) UpdateUser(user *User) (*User, error) {
	result := flywayRepository.DB.Save(user)
	return user, result.Error
}

func (flywayRepository *flywayRepository) DeleteUser(user *User) error {
	result := flywayRepository.DB.Delete(user)
	return result.Error
}
