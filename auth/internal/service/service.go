package service

import (
	"auth/internal/grpc/grpccl"
	"context"
	pbuser "github.com/zxcMentor/grpcproto/protos/user/gen/go"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthService struct {
	clientUser *grpccl.ClientUser
}

func (a *AuthService) Register(email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	mess, err := a.clientUser.CallCreateUser(context.Background(), &pbuser.CreateUserRequest{
		Email:        email,
		HashPassword: string(hashedPassword),
	})
	if err != nil {
		log.Fatal("err call user")
	}

	return mess.Message, nil
}

func (a *AuthService) Login(email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	user,err:=a.clientUser.CallProfileUser(&pbuser.ProfileUserRequest{Id: })
}

func (a *AuthService) ItsValid(token string) (bool, error) {
	return false, nil
}
