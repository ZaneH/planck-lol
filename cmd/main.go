package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"planck-lol/internal/shortener"
)

func main() {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	storage := &shortener.LinkStorage{
		DBPool: dbpool,
	}
	linkService := &shortener.LinkService{
		Storage: storage,
	}

	server := shortener.SetupServer(linkService)
	server.ListenAndServe()
}
