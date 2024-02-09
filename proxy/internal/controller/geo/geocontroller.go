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
	"time"
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

	cacheKey := fmt.Sprintf("geoSearch: %s", req.Input)
	data, err := h.redisClient.Get(context.Background(), cacheKey).Result()
	if err == redis.Nil {

		address, err := h.grpcClient.CallSearchAddress(context.Background(), req)
		if err != nil {
			http.Error(w, "err Call GRPC", http.StatusInternalServerError)
			return
		}

		var addrcache []byte
		addrcache = address.Data
		if err != nil {
			log.Println("adrBts err marsh")
		}
		h.redisClient.Set(context.Background(), cacheKey, addrcache, 20*time.Second)

		w.Write(addrcache)
	} else if err != nil {

		http.Error(w, "Cache retrieval error", http.StatusInternalServerError)
	} else {

		w.Write([]byte(data))
	}

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

	cacheKey := fmt.Sprintf("geoGeocode: %s %s", req.Lon, req.Lat)
	data, err := h.redisClient.Get(context.Background(), cacheKey).Result()
	if err == redis.Nil {
		// Данных нет в кеше, выполняем запрос к gRPC сервису
		address, err := h.grpcClient.CallGeocodeAddress(context.Background(), req)
		if err != nil {
			http.Error(w, "err Call GRPC", http.StatusInternalServerError)
			return
		}

		var addrcache []byte
		addrcache = address.Data
		if err != nil {
			log.Println("adrBts err marsh")
		}
		h.redisClient.Set(context.Background(), cacheKey, addrcache, 20*time.Second)

		w.Write(addrcache)
	} else if err != nil {

		http.Error(w, "Cache retrieval error", http.StatusInternalServerError)
	} else {

		w.Write([]byte(data))
	}
}
