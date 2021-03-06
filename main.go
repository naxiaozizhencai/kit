package main

import (
	"code/Endpoint"
	"code/Service"
	"code/Transport"
	httptranport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func main()  {

	user := Service.UserService{}
	endp := Endpoint.GetUserEndpoint(&user)
	serverHander := httptranport.NewServer(endp, Transport.DecodeUserRequest, Transport.EncodeUserReponse)
	r := mux.NewRouter()
	r.Handle(`/user/{uid:\d+}`, serverHander)
	http.ListenAndServe(":8080", r)
}
