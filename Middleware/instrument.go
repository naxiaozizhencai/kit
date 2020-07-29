package Middleware

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/time/rate"
)

var ErrLimitExceed = errors.New("Rate limit exceed!")
func NewTokenBucketLimitterWithBuildIn(bkt *rate.Limiter) endpoint.Middleware{
	return func(next endpoint.Endpoint) endpoint.Endpoint{
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if bkt.Allow() {
				fmt.Println("aaaaaaaaaaaaaaaaa")
				return nil, ErrLimitExceed
			}
			return next(ctx, request)

		}
	}
}