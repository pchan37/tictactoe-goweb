package server

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/PGo-Projects/output"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/lpar/gzipped"
	"github.com/pchan37/tictactoe-goweb/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const SOCK = "/tmp/tictactoe-goweb.sock"

func withIndexAsDefault(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			r.URL.Path += "index.html"
		}
		h.ServeHTTP(w, r)
	})
}

func Run(cmd *cobra.Command, arg []string) {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)

	webAssetsDirectory := http.Dir(viper.GetString(config.WebAssetsPathKey))
	mux.Method(http.MethodGet, "/*",
		withIndexAsDefault(gzipped.FileServer(webAssetsDirectory)))

	if config.DevRun {
		output.Println("Attempting to run on localhost:8080...", output.BLUE)
		log.Fatal(http.ListenAndServe(":8080", mux))
	} else {
		os.Remove(SOCK)
		unixListener, err := net.Listen("unix", SOCK)
		if err != nil {
			output.Error(err)
		}
		os.Chmod(SOCK, 0666)
		defer unixListener.Close()

		output.Println("Attempting to run on unix socket...", output.BLUE)
		log.Fatal(http.Serve(unixListener, mux))
	}
}
