package merge

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strings"
)

var field map[string]string

func ToString(name string) string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("type %s struct {\n", name))
	keys := make([]string, 0, len(field))
	for key := range field {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		buffer.WriteString(fmt.Sprintf("%s %s\n", key, field[key]))
	}
	buffer.WriteString("}\n")
	return buffer.String()
}

var InvalidStructError = errors.New("string is not a valid struct")

func Resolve(str string) (string, error) {
	str = strings.TrimSpace(str)
	splits := strings.Fields(str)
	if splits[0] != "type" {
		return "", InvalidStructError
	}
	name := splits[1]
	for index := 4; index < len(splits)-1; index = index + 3 {
		if index+3 >= len(splits) {
			return "", InvalidStructError
		}
		if _, ok := field[splits[index]]; !ok {
			field[splits[index]] = splits[index+1] + " " + splits[index+2]
		}
	}
	return name, nil
}
func MergeStructs(items ...string) (string, error) {
	field = make(map[string]string)
	if len(items) < 2 {
		return "", errors.New("should have at least two items")
	}
	var name string
	for _, item := range items {
		n, err := Resolve(item)
		if err != nil {
			return "", err
		}
		name = n
	}
	return ToString(name), nil
}

func SplitStructStrings(reader io.Reader) []string {
	byts, _ := ioutil.ReadAll(reader)
	str := strings.TrimSpace(string(byts))
	splits := strings.Split(str, "}")
	for index := 0; index < len(splits); index++ {
		splits[index] += "}"
	}
	return splits[:len(splits)-1]
}
