package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

func main() {
	_ = godotenv.Load()
	if len(os.Args) < 2 {
		fmt.Println("Error read argument: please run with `artifacts-cleaner $owner/$repo`.")
		os.Exit(1)
	}
	target := os.Args[1]
	split := strings.Split(target, "/")
	if len(split) < 2 {
		fmt.Println("Error read argument: please specify full name for repository like `microsoft/typescript`.")
		os.Exit(1)
	}
	token := os.Getenv("GH_TOKEN")
	if token == "" {
		fmt.Println("Error read access token: please set env `GH_TOKEN` with github access token.")
		os.Exit(1)
	}
	cleaner := NewCleaner(token)
	if err := cleaner.Clean(split[0], split[1]); err != nil {
		fmt.Printf("Error to clean: %s.\n", err.Error())
		os.Exit(1)
	}
}
