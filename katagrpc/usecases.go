package katagrpc

import (
	"context"
	"errors"
	"strings"

	"github.com/farislr/core-proto/kata"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type service struct {
	logger log.Logger
}

func NewService(logger log.Logger) kata.StringServiceServer {
	return &service{
		logger,
	}
}

var errorEmpty = errors.New("Error Empty")

func (svc service) Uppercase(ctx context.Context, request *kata.UppercaseRequest) (*kata.UppercaseResponse, error) {
	logger := log.With(svc.logger, "method", "Uppercase")
	s := request.GetS()

	if s == "" {
		level.Error(logger).Log("err", errorEmpty)
		return &kata.UppercaseResponse{V: ""}, errorEmpty
	}

	logger.Log("uppercase", s)

	return &kata.UppercaseResponse{V: strings.ToUpper(s)}, nil

}
