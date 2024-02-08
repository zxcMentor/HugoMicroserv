package service

import (
	"fmt"

	"log"
	"user/internal/models"
	"user/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepo
}

func NewUserService(repo *repository.UserRepo) *UserService {
	return &UserService{repo}
}

func (u *UserService) CreateUser(email, hashepassword string) (string, error) {
	err := u.Repo.CreateUser(email, hashepassword)
	if err != nil {
		log.Println("err:", err)
		return "", err
	}
	return fmt.Sprint("user created successfully"), nil
}

func (u *UserService) ProfileUser(id int32) (*models.UserDTO, error) {
	user, err := u.Repo.ProfileUser(id)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return user, nil
}

func (u *UserService) ListUsers() ([]*models.UserDTO, error) {
	users, err := u.Repo.ListUsers()
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return users, nil
}
