package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lpar/gzipped"
)

func WithIndexAsDefault(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path += "index.html"
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	staticAssetPath := http.Dir("web/dist")
	mux.Method(http.MethodGet, "/*", WithIndexAsDefault(gzipped.FileServer(staticAssetPath)))

	log.Fatal(http.ListenAndServe(":8080", mux))
}
