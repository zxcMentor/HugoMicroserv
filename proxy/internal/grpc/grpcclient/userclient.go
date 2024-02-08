package grpcclient

import (
	"context"
	pbuser "github.com/zxcMentor/grpcproto/protos/user/gen/go"
	"google.golang.org/grpc"
	"log"
)

type ClientUser struct{}

func NewClientUser() *ClientUser {
	return &ClientUser{}
}

func (c *ClientUser) CallProfileUser(ctx context.Context, req *pbuser.ProfileUserRequest) (*pbuser.ProfileUserResponse, error) {

	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := pbuser.NewUserServiceClient(conn)
	res, err := client.ProfileUser(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}

func (c *ClientUser) CallListUsers(ctx context.Context, req *pbuser.ListUsersRequest) (*pbuser.ListUsersResponse, error) {

	conn, err := grpc.Dial("user:50053", grpc.WithInsecure())
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	client := pbuser.NewUserServiceClient(conn)
	res, err := client.ListUsers(ctx, req)
	if err != nil {
		log.Println("err:", err)
		return nil, err
	}
	return res, nil
}
