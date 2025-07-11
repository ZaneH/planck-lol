package shortener

import (
	"fmt"
	"net/http"
	"os"
)

func SetupServer(s *LinkService) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", handleCodeCreate(s))
	mux.HandleFunc("/{code}", handleCodeGet(s))
	mux.HandleFunc("/", handleIndex)

	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return server
}
