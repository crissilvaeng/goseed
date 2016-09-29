package api

import (
	"fmt"
	"net/http"
)

// Routes define the handled endpoint on API.
func Routes() {
	http.HandleFunc("/", hello)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, World!")
}
