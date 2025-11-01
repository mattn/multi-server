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

const version = "0.0.3"

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

		http.FileServer(http.Dir(siteDir)).ServeHTTP(w, r)
	})

	logger := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print(r.RemoteAddr + " " + r.Method + " " + r.URL.String())
		http.DefaultServeMux.ServeHTTP(w, r)
	})
	log.Println("Server starting on " + addr)
	log.Fatal(http.ListenAndServe(addr, logger))
}
