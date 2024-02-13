package main

import (
	"github.com/mjmcconnell/feature-manager/internal"
	// modules
	_ "github.com/mjmcconnell/feature-manager/modules/hello"
)

func main() {
	router := internal.NewRouter()
	router.Start("0.0.0.0:8080")
}
