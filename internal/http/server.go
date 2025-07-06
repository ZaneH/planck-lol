package internal

import (
	"net/http"
	"planck-lol/internal/controller"
)

func SetupServer(c *controller.LinkController) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/create", CreateCodeHandle(c))
	mux.HandleFunc("/{code}", GetCodeHandle(c))
	mux.HandleFunc("/", IndexHandle)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	return server
}
