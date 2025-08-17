package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserByEmail(email string) (domain.User, error)
	FindUserById(id uint) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(u domain.User) (domain.User, error) {
	err := r.db.Create(&u).Error
	if err != nil {
		log.Printf("create user error: %v", err)
		return domain.User{}, errors.New("failed to create user")
	}

	return u, nil
}

func (r userRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User

	err := r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Printf("find user by email error: %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User

	err := r.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("find user by id error: %v", err)
		return domain.User{}, errors.New("user does not exist")
	}
	return user, nil
}

func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User

	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id = ?", id).Updates(u).Error
	if err != nil {
		log.Printf("update user error: %v", err)
		return domain.User{}, errors.New("failed to update user")
	}

	return user, nil
}
