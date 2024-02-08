package user

import (
	"context"
	"encoding/json"
	pbuser "github.com/zxcMentor/grpcproto/protos/user/gen/go"
	"log"
	"net/http"
	"proxy/internal/grpc/grpcclient"
	"proxy/internal/models"
)

type HandleUser struct {
	clientuser *grpcclient.ClientUser
}

func NewHandleUser(clUser *grpcclient.ClientUser) *HandleUser {
	return &HandleUser{clUser}
}

func (h *HandleUser) ProfileUser(w http.ResponseWriter, r *http.Request) {
	email := "@example"
	req := &pbuser.ProfileUserRequest{Email: email}

	res, err := h.clientuser.CallProfileUser(context.Background(), req)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	user := &models.User{
		Id:    res.User.Id,
		Email: res.User.Email,
	}
	jsData, err := json.Marshal(user)
	if err != nil {
		log.Println("err:", err)
	}
	w.Write(jsData)
}

func (h *HandleUser) ListUsers(w http.ResponseWriter, r *http.Request) {

	req := &pbuser.ListUsersRequest{}
	res, err := h.clientuser.CallListUsers(context.Background(), req)
	if err != nil {
		log.Println("err:", err)
		http.Error(w, "err serv", http.StatusInternalServerError)
		return
	}
	var users []models.User
	for _, v := range res.Users {
		user := models.User{
			Id:    v.Id,
			Email: v.Email,
		}
		users = append(users, user)
	}
	jsData, err := json.Marshal(users)
	if err != nil {
		log.Println("err :", err)
	}
	w.Write(jsData)
}
