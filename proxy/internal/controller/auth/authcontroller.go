package auth

import (
	"context"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
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

	req := &pbauth.RegisterRequest{
		Email:    "@example",
		Password: "1234",
	}
	mess, err := h.clientauth.CallRegister(context.Background(), req)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(mess.Message))

}

func (h *HandleAuth) Login(w http.ResponseWriter, r *http.Request) {
	req := &pbauth.LoginRequest{
		Email:    "@example",
		Password: "1234",
	}
	token, err := h.clientauth.CallLogin(context.Background(), req)
	if err != nil {
		http.Error(w, "err register failed", http.StatusInternalServerError)
		return
	}

	w.Write([]byte(token.Token))
}
