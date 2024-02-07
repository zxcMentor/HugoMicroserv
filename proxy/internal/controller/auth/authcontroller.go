package auth

/*
import (
	"log"
	"net/http"
	"proxy/internal/grpc/grpcclient"
)

type HandleAuth struct {
	clientauth *grpcclient.ClientAuth
}

func NewHandleAuth() *HandleAuth {
	return &HandleAuth{clientauth: &grpcclient.ClientAuth{}}
}

func (h *HandleAuth) Register(w http.ResponseWriter, r *http.Request) {

		req := &pbauth.RegisterRequest{
			Email:    "@example",
			Password: "1234",
		}
		mess, err := h.clientauth.CallRegister(context.Background(), req)
		if err != nil {
			http.Error(w, "err register failed", http.StatusInternalServerError)
			return
		}
		log.Println("user registered")

		w.Write([]byte(mess.Message))



}

func (h *HandleAuth) Login(w http.ResponseWriter, r *http.Request) {

}

*/
