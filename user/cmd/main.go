package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"log"
	"microservice/user/internal/grpc/user"
	"microservice/user/internal/repository"
	"microservice/user/internal/service"
	pbuser "microservice/user/protos/gen/go"
	"net"
)

func main() {
	// Подключение к базе данных
	db, err := sqlx.Connect("postgres", "user=username dbname=database sslmode=disable")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) UNIQUE NOT NULL,
        hashed_password VARCHAR(255) NOT NULL
    )`)

	if err != nil {
		panic(err)
	}

	userRepo := repository.NewUserRepo(db)

	userService := service.NewUserService(userRepo)

	serviceUser := user.NewServiceUser(userService)

	// Создание gRPC сервера
	lis, err := net.Listen("tcp", "50053")
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", "50053", err)
	}
	grpcServer := grpc.NewServer()

	// Регистрация ServiceUser в gRPC сервере
	pbuser.RegisterUserServiceServer(grpcServer, serviceUser)

	// Запуск gRPC сервера
	log.Printf("Starting gRPC server on port %s", "50053")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
