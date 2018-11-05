package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/golang/glog"

	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/generate"
)

func main() {
	flag.Parse()
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
	err = generate.GenerateServiceDoc(myapi)
	if err != nil {
		glog.Fatal(err)
	}
}
