package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"planck-lol/internal/shortener"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func main() {
	err := run()
	if err != nil {
		log.Fatalf("Unhandled error: %v\n", err)
	}
}

func run() error {
	pgurl := os.Getenv("DATABASE_URL")
	dbpool, err := pgxpool.New(context.Background(), pgurl)

	if err != nil {
		return err
	}

	rurl := os.Getenv("REDIS_URL")
	opts, err := redis.ParseURL(rurl)
	rdb := redis.NewClient(opts)

	if err != nil {
		return err
	}

	defer dbpool.Close()

	storage := &shortener.LinkStorage{DBPool: dbpool, Cache: rdb}
	linkService := &shortener.LinkService{Storage: storage}
	server := shortener.SetupServer(linkService)

	fmt.Printf("Listening on %s\n", server.Addr)
	server.ListenAndServe()
	return nil
}
