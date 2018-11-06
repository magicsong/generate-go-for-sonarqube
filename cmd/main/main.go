package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"

	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/generate"
)

func main() {
	flag.Parse()
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	file, err := os.Open("api.json")
	if err != nil {
		glog.Fatal(err)
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	myapi := new(api.API)
	err = decoder.Decode(myapi)
	if err != nil {
		glog.Fatal(err)
	}
	err = generate.Build(myapi)
	if err != nil {
		glog.Fatal(err)
	}
}
