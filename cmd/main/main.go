package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/Mishany52/testTaskSWoyo/internal/config"
	url2 "github.com/Mishany52/testTaskSWoyo/internal/url"
	url "github.com/Mishany52/testTaskSWoyo/internal/url/db"

	"github.com/Mishany52/testTaskSWoyo/pkg/client/postgresql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)
func main(){

	var repository url2.Repository
	isPostgres := checkRunKey();
	
	router := chi.NewRouter()

	cfg := config.GetConfig()

	if isPostgres {
		client, err := postgresql.NewClient(context.TODO(),cfg.PostgresURL)
		if err != nil {
			log.Fatalf("%v", err)
		}
		repository = url.NewRepository(client)
		
	} else {
		repository = url.NewRepositoryMap()
	}

	handler := url2.NewHandler(repository)

	router.Use(middleware.Logger)
	handler.Register(router)

	start(router, cfg)
}

func start(router *chi.Mux, cfg *config.Config)  {
	// var ud = &db.UrlDb{}
	// ud.InitializeDB()

	listener, err := net.Listen("tcp", cfg.ServerAddr)
	
	if err != nil {
		panic(err)
	}
	server := &http.Server{
		Handler: router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}
	
	fmt.Printf("App run on: %s \n", cfg.ServerAddr)
	log.Fatalln(server.Serve(listener))
}

func checkRunKey () bool {
	if (len(os.Args) > 2) {
		panic("usage go run main.go . [-d]")
	}
	if len(os.Args) == 1 {
		return false
	} else if os.Args[1] == "-d" {
		return true
	} else {
		panic("usage go run main.go [-d] or go run main.go")
	}
}