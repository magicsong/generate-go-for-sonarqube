package main

import (
	"log"
	"os"

	"github.com/magicsong/generate-go-for-sonarqube/pkg/static"
)

func main() {
	err := static.GenerateStaticFiles(os.Args[1], os.Args[2])
	if err != nil {
		log.Fatal(err)
	}
}
