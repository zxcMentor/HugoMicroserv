package grpcclient

import (
	"context"
	pbgeo "github.com/zxcMentor/protos/grpcproto/geo/protos/gen/go"
	"google.golang.org/grpc"
	"log"
)

type ClientGeo struct{}

func NewClientGeo() *ClientGeo {
	return &ClientGeo{}
}

func (c *ClientGeo) CallSearchAddress(ctx context.Context, req *pbgeo.SearchRequest) (*pbgeo.SearchResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pbgeo.NewGeoServiceClient(conn)

	res, err := client.SearchAddress(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове gRPC: %v", err)
		return nil, err
	}

	return res, nil
}

func (c *ClientGeo) CallGeocodeAddress(ctx context.Context, req *pbgeo.GeocodeRequest) (*pbgeo.GeocodeResponse, error) {
	conn, err := grpc.Dial("geo:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Ошибка при подключении к серверу: %v", err)
		return nil, err
	}
	defer conn.Close()
	client := pbgeo.NewGeoServiceClient(conn)

	res, err := client.GeocodeAddress(context.Background(), req)
	if err != nil {
		log.Fatalf("Ошибка при вызове gRPC: %v", err)
		return nil, err
	}

	return res, nil
}
