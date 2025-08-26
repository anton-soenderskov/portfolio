package server

import (
	"net/http"
)

func CreateServer() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	})

	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir("./pages/about.html")).ServeHTTP(w, r)
	})

	server := http.Server{}
	server.Handler = mux
	server.Addr = ":8080"
	return server.ListenAndServe()
}
