package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	var src http.Server
	// http2.VerboseLogs = true
	src.Addr = ":8080"
	// http2.ConfigureServer(&src, nil)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi testerer %q \n", html.EscapeString(r.URL.Path))
		ShowRequestInfoHandler(w, r)
	})
	log.Fatal(src.ListenAndServeTLS("server.crt", "server.key"))
}

func ShowRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RemoteURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d\n", r.ContentLength)
	fmt.Fprintf(w, "TLS: %s\n", r.TLS)
	fmt.Fprintf(w, "\n Headers:")
	r.Header.Write(w)
}
