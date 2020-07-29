package main

import (
	"code/Endpoint"
	"code/Middleware"
	"code/Service"
	"code/Transport"
	httptranport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"golang.org/x/time/rate"
	"net/http"
	"time"
)

func main()  {

	var svc Service.IUserService
	svc = &Service.UserService{}

	svc = Middleware.LoggingMiddleware()(svc)
	ratebucket := rate.NewLimiter(rate.Every(time.Second*1), 3)

	endp := Endpoint.GetUserEndpoint(svc)
	endp = Middleware.NewTokenBucketLimitterWithBuildIn(ratebucket)(endp)
	serverHander := httptranport.NewServer(endp, Transport.DecodeUserRequest, Transport.EncodeUserReponse)
	r := mux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, serverHander)
	http.ListenAndServe(":8080", r)
}
