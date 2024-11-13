package repository

import (
    "gorm.io/gorm"
    "github.com/pp00x/hydrate/internal/model"
)

type UserRepository interface {
    Create(user *model.User) error
    FindByEmail(email string) (*model.User, error)
    FindByID(id uint) (*model.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db}
}

func (r *userRepository) Create(user *model.User) error {
    return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
    var user model.User
    result := r.db.Where("email = ?", email).First(&user)
    return &user, result.Error
}

func (r *userRepository) FindByID(id uint) (*model.User, error) {
    var user model.User
    result := r.db.First(&user, id)
    return &user, result.Error
}