package main

import (
	"fmt"
	"net/http"

	"go-api-service/internal/controllers"
)

func main() {
	controller := controllers.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/", controller.HandleRequests)

	server := &http.Server{
		Addr: fmt.Sprintf("localhost:5000"),
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
