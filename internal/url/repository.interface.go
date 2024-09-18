package url

import (
	"context"
)

type Repository interface {
	Create(ctx context.Context, longUrl, shortUrl string) error
	FindOneLongByShort(ctx context.Context, shortUrl string) (string,error)
	FindOneShortByLong(ctx context.Context, longUrl string) (string,error)
}
