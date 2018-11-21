package main

import (
	"fmt"
	"os"

	. "github.com/magicsong/generate-go-for-sonarqube/pkg/merge"
)

func main() {
	file, err := os.OpenFile(os.Args[1], os.O_RDWR, os.ModeAppend)
	defer file.Close()
	errorHandler(err)
	splits := SplitStructStrings(file)
	str, err := MergeStructs(splits...)
	errorHandler(err)
	_, err = file.WriteString("\n\n-------Generated code-------\n" + str)
	errorHandler(err)
}
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
