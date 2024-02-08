package geo

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	pbgeo "github.com/zxcMentor/grpcproto/protos/geo/gen/go"
	"log"
	"net/http"
	"proxy/internal/grpc/grpcclient"
)

type HandleGeo struct {
	grpcClient  *grpcclient.ClientGeo
	redisClient *redis.Client
}

func NewHandGeo(clientGeo *grpcclient.ClientGeo) *HandleGeo {
	redcl := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})

	return &HandleGeo{
		grpcClient:  clientGeo,
		redisClient: redcl,
	}
}

func (h *HandleGeo) SearchHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("sear run")
	req := &pbgeo.SearchRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	/*
		cachekey := "geo:Search " + req.Input
		result, err := h.redisClient.Get(context.Background(), cachekey).Result()
		if err == redis.Nil {

			address, err := h.grpcClient.CallSearchAddress(context.Background(), req)
			if err != nil {
				http.Error(w, "err grpc response", http.StatusInternalServerError)
				return
			}

			err = h.redisClient.Set(context.Background(), cachekey, address.Data, time.Hour).Err()
			if err != nil {
				http.Error(w, "REDIS SET Internal Server Error", http.StatusInternalServerError)
				return
			}

			result = string(address.Data)
		} else if err != nil {
			log.Println("err REDIS", err)
			return
		}

		var cachedAddress pbgeo.SearchResponse
		err = json.Unmarshal([]byte(result), &cachedAddress)
		if err != nil {
			http.Error(w, "err Unmarshal", 0)
			return
		}

		var res []byte
		res = cachedAddress.Data

	*/

	var resp []byte
	address, err := h.grpcClient.CallSearchAddress(context.Background(), req)
	if err != nil {
		http.Error(w, "err Call GRPC", http.StatusBadRequest)
	}
	resp = address.Data
	w.Write(resp)

}

func (h *HandleGeo) GeocodeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("geocode run")
	req := &pbgeo.GeocodeRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("err read body")
		http.Error(w, "Invalid request format", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var resp []byte
	address, err := h.grpcClient.CallGeocodeAddress(context.Background(), req)
	if err != nil {
		http.Error(w, "err Call GRPC", http.StatusBadRequest)
	}
	resp = address.Data
	w.Write(resp)
}
