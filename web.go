package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)

	port := os.Getenv("PORT")

	if port != "" {
		fmt.Println("Listining in 0.0.0.0:" + port)
	} else {
		fmt.Println("PORT environment variable is not found, using the default value to them instead.")
		port = "80"
		fmt.Println("Listining in 0.0.0.0:" + port)
	}

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, World!")
}
