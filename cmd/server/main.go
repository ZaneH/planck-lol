package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"os"
	"planck-lol/internal/shortener"
)

func main() {
	err := run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unhandled error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return err
	}

	defer dbpool.Close()

	storage := &shortener.LinkStorage{DBPool: dbpool}
	linkService := &shortener.LinkService{Storage: storage}
	server := shortener.SetupServer(linkService)

	fmt.Printf("Listening on %s\n", server.Addr)
	server.ListenAndServe()
	return nil
}
