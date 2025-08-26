package server

import (
	"net/http"
	"path/filepath"
)

func CreateServer() error {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", serveHTML("index.html"))
	mux.HandleFunc("/about", serveHTML("pages/about.html"))
	mux.HandleFunc("/projects", serveHTML("pages/projects.html"))
	mux.HandleFunc("/blog", serveHTML("pages/blog.html"))

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	return server.ListenAndServe()
}

func serveHTML(filename string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		http.ServeFile(w, r, filepath.Clean(filename))
	}
}
