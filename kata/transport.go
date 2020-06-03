package kata

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	Uppercase endpoint.Endpoint
	Count     endpoint.Endpoint
}

func NewTransport(s StringService) Endpoints {
	return Endpoints{
		Uppercase: makeUppercaseEndpoint(s),
		Count:     makeCountEndpoint(s),
	}
}

func makeUppercaseEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UppercaseRequest)

		v, err := svc.Uppercase(ctx, req.S)
		if err != nil {
			return UppercaseResponse{v, err.Error()}, nil
		}

		return UppercaseResponse{v, ""}, nil
	}
}

func makeCountEndpoint(svc StringService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CountRequest)
		v := svc.Count(ctx, req.S)
		return CountResponse{v}, nil
	}
}
