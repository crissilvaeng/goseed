package main

import (
	"fmt"
	"os"

	"github.com/crissilvaeng/goseed/api"
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

	if port == "" {
		return
	}

	api.Run(port)
}
