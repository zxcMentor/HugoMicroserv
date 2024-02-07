package main

import (
	"geo/internal/grpc/geo"
	pbgeo "geo/protos/gen/go"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	pbgeo.RegisterGeoServiceServer(server, &geo.ServerGeo{})

	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
