package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/crissilvaeng/goddamned/api"
)

// Provides by govvv in compile-time.
var (
	// GitBranch is current branch name the code is built off.
	GitBranch string
	// GitSummary is output of git describe --tags --dirty --always.
	GitSummary string
	// BuildDate is RFC3339 formatted UTC date in compile moment.
	BuildDate string
)

func main() {
	fmt.Printf("build=%s-%s date=%s\n\n", GitBranch, GitSummary, BuildDate)

	port := os.Getenv("PORT")

	if port != "" {
		fmt.Println("Listining in 0.0.0.0:" + port)
	} else {
		fmt.Println("PORT environment variable is not found, using the default value to them instead.")
		port = "80"
		fmt.Println("Listining in 0.0.0.0:" + port)
	}

	api.Routes()

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}
