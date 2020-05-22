package kata

import (
	"context"
	"errors"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

type stringService struct {
	logger log.Logger
}

func NewService(logger log.Logger) StringService {
	return &stringService{
		logger,
	}
}

var errorEmpty = errors.New("Error Empty")

func (svc stringService) Uppercase(ctx context.Context, s string) (string, error) {
	logger := log.With(svc.logger, "method", "Uppercase")

	if s == "" {
		level.Error(logger).Log("err", errorEmpty)
		return "", errorEmpty
	}

	logger.Log("uppercase", s)

	return strings.ToUpper(s), nil
}

func (svc stringService) Count(ctx context.Context, s string) int {
	logger := log.With(svc.logger, "method", "Count")

	logger.Log("count", s)

	return len(s)
}
