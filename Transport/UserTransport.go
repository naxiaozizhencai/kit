package Transport

import (
	"code/Endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request)(interface{}, error){
	if r.URL.Query().Get("uid") != "" {
		uid, _ := strconv.Atoi(r.URL.Query().Get("uid"))
		return Endpoint.UserRequest{
			Uid:uid,
		}, nil
	}

	return nil, errors.New("参数错误")
}

func EncodeUserReponse(c context.Context, w http.ResponseWriter,reponse interface{})(error){
	w.Header().Set("content-type", "application/json")
	return json.NewEncoder(w).Encode(reponse)
}
