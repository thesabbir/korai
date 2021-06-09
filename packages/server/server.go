package server

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
)

func server() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Mount("/api", ApiRouter())

	addr := ":9000"
	fmt.Printf("Starting server on %v\n", addr)

	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal("Error! Failed to start server.")
	}
}

func Start() {
	server()
}
