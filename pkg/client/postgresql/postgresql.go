package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/Mishany52/testTaskSWoyo/internal/config"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	_ "github.com/lib/pq"
)

var conf = config.GetConfig()

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, urlConnection string) (conn *pgx.Conn, err error) {
	fmt.Printf("urlSql: %s \n",urlConnection)
	conn, err = pgx.Connect(ctx, urlConnection)
	if err != nil {
		log.Fatal("error postgresql connection")
	}
	return conn, nil
}


// func (td *UrlDb) createConnection() {
	
// 	//Открываем соединение
// 	db, err := sql.Open("postgres", conf.PostgresURL)

// 	if err != nil {
// 		panic(err)
// 	}

// 	//Проверка соединения
// 	err = db.Ping()

// 	if err != nil {
// 		panic(err)
// 	}

// 	log.Println("Successfully connected!")

// 	td.conn = db
// }

