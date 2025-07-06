package internal

import (
	"fmt"
	"net/http"
	"os"
	"planck-lol/internal/controller"
)

func SetupServer(c *controller.LinkController) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", CreateCodeHandle(c))
	mux.HandleFunc("/{code}", GetCodeHandle(c))
	mux.HandleFunc("/", IndexHandle)

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
