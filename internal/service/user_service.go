package service

import (
    "errors"

    "github.com/pp00x/hydrate/internal/model"
    "github.com/pp00x/hydrate/internal/repository"
)

type UserService interface {
    Register(user *model.User) error
    Login(email, password string) (*model.User, error)
    GetUserByID(id uint) (*model.User, error)
}

type userService struct {
    userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
    return &userService{userRepo}
}

func (s *userService) Register(user *model.User) error {
    hashedPassword, err := HashPassword(user.Password)
    if err != nil {
        return err
    }
    user.Password = hashedPassword
    return s.userRepo.Create(user)
}

func (s *userService) Login(email, password string) (*model.User, error) {
    user, err := s.userRepo.FindByEmail(email)
    if err != nil {
        return nil, errors.New("invalid email or password")
    }
    if !CheckPasswordHash(password, user.Password) {
        return nil, errors.New("invalid email or password")
    }
    return user, nil
}

func (s *userService) GetUserByID(id uint) (*model.User, error) {
    return s.userRepo.FindByID(id)
}