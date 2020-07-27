package Middleware

import (
	"code/Service"
	"log"
)
type ServerMiddleware func(Service.IUserService)Service.IUserService

type loggingsMiddleware struct {
	Service.IUserService
}

func LoggingMiddleware() ServerMiddleware    {
	return func(next Service.IUserService) Service.IUserService {
		return loggingsMiddleware{next}
	}
}

func (mw loggingsMiddleware) GetName(userId int) string {
	defer log.Println("aaaaaaaaaaaaaaaaaaaa")
	defer log.Println("bbbbbbbbbbbbb")
	return mw.IUserService.GetName(userId)
}