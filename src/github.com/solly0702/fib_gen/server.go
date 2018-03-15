package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/solly0702/fib_gen/api"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/api/fib-gen", api.FibHandleFunc)

	http.ListenAndServe(port(), nil)
}

func port() string {
	port := os.Getenv("PORT")

	if len(port) == 0 {
		port = "5000"
	}
	return ":" + port
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "<h1>Welcome to Fibonacci Generator!</h1>")
}
