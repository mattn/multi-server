package main

import (
	"flag"
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
	var addr string
	var sitesDir string
	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.StringVar(&sitesDir, "sites-dir", "/data", "Directory containing site folders")
	flag.Parse()

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
		fs.ServeHTTP(w, r)
	})

	log.Println("Server starting on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
