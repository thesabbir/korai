package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/thesabbir/korai/packages/sysinfo"
	"log"
	"net/http"
)


func ApiRouter() http.Handler {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("welcome"))
		if err != nil {
			log.Fatal("Error welcome")
		}
	})
	r.Get("/sysinfo", func(w http.ResponseWriter, r *http.Request) {
		info, _ := sysinfo.Info()
		_, err := w.Write(info)
		if err != nil {
			log.Fatal("Error welcome")
		}
	})
	return r
}
