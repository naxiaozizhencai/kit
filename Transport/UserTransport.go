package Transport

import (
	"code/Endpoint"
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func DecodeUserRequest(c context.Context, r *http.Request)(interface{}, error){
	vars := mux.Vars(r)
	if uid, ok := vars["uid"];ok{
		uid, _ := strconv.Atoi(uid)
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
