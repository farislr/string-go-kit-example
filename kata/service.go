package kata

import "context"

type StringService interface {
	Uppercase(ctx context.Context, str string) (string, error)
	Count(ctx context.Context, str string) int
}
