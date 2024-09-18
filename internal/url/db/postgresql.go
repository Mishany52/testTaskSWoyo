package url

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Mishany52/testTaskSWoyo/internal/url"
	"github.com/Mishany52/testTaskSWoyo/pkg/client/postgresql"
	"github.com/jackc/pgx/v5/pgconn"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) url.Repository {
	return &repository{
		client: client,
	}
}

func formatQuery(q string) string {
	return strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " ")
}

func (r *repository) Create(ctx context.Context, longUrl, shortUrl string) error{
	q := `
			INSERT INTO url 
				(longUrl, shortUrl)
			VALUES 
				($1, $2)
			RETURNING id;
	`
	fmt.Printf("Sql Query: %s \n", formatQuery(q))
	var id int64
	if err := r.client.QueryRow(ctx, q, longUrl, shortUrl).Scan(&id); err!=nil{
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr){
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Println(newErr)
			return newErr
		}
	}
	return nil
}

func (r *repository) FindOneLongByShort(ctx context.Context, shortUrl string) (longUrl string, err error){
	q :=`
		SELECT longUrl FROM public.url WHERE shortUrl = $1;
	`
	fmt.Printf("Sql Query: %s \n", formatQuery(q))
	err = r.client.QueryRow(ctx, q, shortUrl).Scan(&longUrl)
	if err != nil {
		return "", nil
	}
	return longUrl, err
}
func (r *repository) FindOneShortByLong(ctx context.Context, longUrl string) (shortUrl string, err error){
	q :=`
		SELECT shortUrl FROM public.url WHERE longUrl = $1;
	`
	fmt.Printf("Sql Query: %s \n", formatQuery(q))
	err = r.client.QueryRow(ctx, q, longUrl).Scan(&shortUrl)
	if err != nil {
		return "", nil
	}
	return shortUrl, err
}

