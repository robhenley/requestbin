package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	var port string
	var host string

	flag.StringVar(&host, "host", "127.0.0.1", "The host to listen on.")
	flag.StringVar(&port, "port", "8080", "The port to listen on.")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

		req := bytes.Buffer{}
		r.Write(&req)

		fmt.Printf("\n=== %s ===\n%s\n",
			time.Now().Format("15:04:05"),
			req.String(),
		)
	})

	fmt.Printf("Listening for requests at %s:%s\n", host, port)
	http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), nil)
}
