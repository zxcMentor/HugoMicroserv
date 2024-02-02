package geo

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	pb "microservice/geo/protos/gen/go"
	"net/http"
	"time"
)

type HandleGeo struct {
	geoClient   pb.GeoServiceClient
	redisClient *redis.Client
}

func (h *HandleGeo) SearchHandle(w http.ResponseWriter, r *http.Request) {
	get := r.Header.Get("Authorization")
	req := &pb.SearchRequest{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		http.Error(w, "err decode json", http.StatusBadRequest)
	}
	defer r.Body.Close()
	cachekey := "geo:Search " + req.Input
	result, err := h.redisClient.Get(context.Background(), cachekey).Result()
	if err == redis.Nil {
		address, err := h.geoClient.SearchAddress(context.Background(), req)
		if err != nil {
			http.Error(w, "err grpc response", http.StatusInternalServerError)
			return
		}

		jsAddress, err := json.Marshal(address)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = h.redisClient.Set(context.Background(), cachekey, jsAddress, time.Hour).Err()
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		result = string(jsAddress)
	} else if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	var cachedAddress pb.SearchResponse
	err = json.Unmarshal([]byte(result), &cachedAddress)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(&cachedAddress)
	if err != nil {
		http.Error(w, "err encode json", http.StatusBadRequest)
		return
	}

}

func (h *HandleGeo) GeocodeHandle(w http.ResponseWriter, r *http.Request) {

}
