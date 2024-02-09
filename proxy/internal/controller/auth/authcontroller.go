package auth

import (
	"context"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"proxy/internal/grpc/grpcclient"
)

type HandleAuth struct {
	clientauth *grpcclient.ClientAuth
}

func NewHandleAuth(clientAuth *grpcclient.ClientAuth) *HandleAuth {
	return &HandleAuth{clientauth: clientAuth}
}

func (h *HandleAuth) Register(w http.ResponseWriter, r *http.Request) {
	email := "@example"
	password := "1234"
	hashepassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("err generate hashepassword")
	}

	req := &pbauth.RegisterRequest{
		Email:         email,
		Hashepassword: string(hashepassword),
	}
	mess, err := h.clientauth.CallRegister(context.Background(), req)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(mess.Message))

}

func (h *HandleAuth) Login(w http.ResponseWriter, r *http.Request) {
	email := "@example"
	password := "1234"

	req := &pbauth.LoginRequest{
		Email:    email,
		Password: password,
	}

	token, err := h.clientauth.CallLogin(context.Background(), req)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Authorization", "Bearer "+token.Token)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token.Token))
}
