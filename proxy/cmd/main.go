package main

import (
	"log"
	"net/http"
	"proxy/internal/controller/geo"
	"proxy/internal/grpc/grpcclient"
	"proxy/internal/router"
)

func main() {
	gcl := grpcclient.NewClientGeo()
	hg := geo.NewHandGeo(gcl)

	r := router.StRout(hg)

	log.Println("proxyq serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}
