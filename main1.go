package main

import (
	"code/Endpoint"
	"code/Middleware"
	"code/Service"
	"code/Transport"
	httptranport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {

	var svc Service.IUserService
	svc = &Service.UserService{}

	svc = Middleware.LoggingMiddleware()(svc)
	endp := Endpoint.GetUserEndpoint(svc)
	serverHander := httptranport.NewServer(endp, Transport.DecodeUserRequest, Transport.EncodeUserReponse)
	r := mux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, serverHander)
	http.ListenAndServe(":8080", r)
}
