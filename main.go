package main

import (
	"code/Endpoint"
	"code/Service"
	"code/Transport"
	httptranport "github.com/go-kit/kit/transport/http"
	"net/http"
)

func main()  {
	user := Service.UserService{}
	endp := Endpoint.GetUserEndpoint(&user)
	serverHander := httptranport.NewServer(endp, Transport.DecodeUserRequest, Transport.EncodeUserReponse)
	http.ListenAndServe(":8080", serverHander)
}
