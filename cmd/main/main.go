package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/golang/glog"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/api"
	"github.com/magicsong/generate-go-for-sonarqube/pkg/generate"
)

var (
	h           bool
	JsonPath    string
	OutputPath  string
	PackageName string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.StringVar(&JsonPath, "f", "", "specify location of api file(only support local file)")
	flag.StringVar(&OutputPath, "o", ".", "specify the destination dir, default is current workspace")
	flag.StringVar(&PackageName, "n", "sonarqube", "specify the name of generated package,default is \"sonarqube\"")
	flag.Usage = usage
}

func main() {
	flag.Parse()
	if h {
		flag.Usage()
		return
	}
	if JsonPath == "" {
		glog.Fatal("Must specify the json location,please add -f [filepath]")
	}
	_, err := os.Stat(JsonPath) //os.Stat获取文件信息
	if err != nil {
		glog.Fatal("No such api file")
	}

	file, err := os.Open(JsonPath)
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
	err = generate.Build(PackageName, OutputPath, myapi)
	if err != nil {
		glog.Fatal(err)
	}
}

func usage() {
	fmt.Fprintf(os.Stderr, ` generate-go-for-sonarqube version: 0.0.1
Usage: main.go [-h] -f jsonpath [-o outputpath] 

Options:
`)
	flag.PrintDefaults()
}
