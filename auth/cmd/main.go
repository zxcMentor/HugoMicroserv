package main

import (
	"auth/internal/grpc/auth"
	pbauth "github.com/zxcMentor/grpcproto/protos/auth/gen/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	pbauth.RegisterAuthServiceServer(server, &auth.ServiceAuth{})

	log.Println("Запуск gRPC сервера auth...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
