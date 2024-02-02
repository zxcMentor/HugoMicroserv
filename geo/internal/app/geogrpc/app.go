package geogrpc

import (
	"google.golang.org/grpc"
	"log"
	"microservice/geo/internal/grpc/geo"
	pb "microservice/geo/protos/gen/go"
	"net"
)

func StartGRPC() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Ошибка при прослушивании порта: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterGeoServiceServer(server, &geo.ServerGeo{})

	log.Println("Запуск gRPC сервера...")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Ошибка при запуске сервера: %v", err)
	}
}
