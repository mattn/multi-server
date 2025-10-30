package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
)

const name = "multi-server"

const version = "0.0.1"

var revision = "HEAD"

func main() {
	const sitesDir = "/data"

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		host := r.Host
		if h, _, err := net.SplitHostPort(host); err == nil {
			host = h
		}

		siteDir := filepath.Join(sitesDir, host)
		log.Printf("Request for host: %s => serving from %s", host, siteDir)

		if _, err := os.Stat(siteDir); os.IsNotExist(err) {
			http.Error(w, "Site not found", http.StatusNotFound)
			return
		}

		fs := http.FileServer(http.Dir(siteDir))
		fs = spaFileServer(fs, siteDir)
		fs.ServeHTTP(w, r)
	})

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func spaFileServer(next http.Handler, root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := filepath.Join(root, r.URL.Path)

		if _, err := os.Stat(path); os.IsNotExist(err) {
			indexPath := filepath.Join(root, "index.html")
			if _, err := os.Stat(indexPath); err == nil {
				http.ServeFile(w, r, indexPath)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
