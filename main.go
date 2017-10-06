package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/freeformz/priv"
)

func private(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, priv.Ate)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	fmt.Println(http.ListenAndServe(":"+port, http.HandlerFunc(private)))
}
