package main

import (
	"log"
	"net/http"
	"proxy/internal/controller/auth"
	"proxy/internal/controller/geo"
	"proxy/internal/controller/user"
	"proxy/internal/grpc/grpcclient"
	"proxy/internal/router"
)

func main() {
	gcl := grpcclient.NewClientGeo()
	acl := grpcclient.NewClientAuth()
	ucl := grpcclient.NewClientUser()
	hg := geo.NewHandGeo(gcl)
	ah := auth.NewHandleAuth(acl)
	uh := user.NewHandleUser(ucl)
	r := router.StRout(hg, ah, uh)

	log.Println("proxyq serv started on ports :8080")
	http.ListenAndServe(":8080", r)
}
