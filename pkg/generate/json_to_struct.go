package generate

import (
	"bytes"
	"errors"
	"strings"

	"github.com/magicsong/generate-go-for-sonarqube/pkg/gojson"
)

var converter = gojson.NewJSONToGOConverter()

func ConvertStringToStruct(json, name string) (string, error) {
	if json == "" {
		return "", errors.New("Json string must no be empty")
	}
	reader := new(bytes.Buffer)
	reader.WriteString(json)
	byts, err := converter.Generate(reader, gojson.ParseJson, name, []string{"json"}, true)
	if err != nil {
		return "", err
	}
	return string(byts), nil
}
func UnionJSONToStruct(jsons []string, name string) (string, error) {
	if len(jsons) == 0 {
		return "", errors.New("Jsons string must no be zero")
	}
	reader := new(bytes.Buffer)
	reader.WriteString("[")
	reader.WriteString(strings.Join(jsons, ","))
	reader.WriteString("]")
	byts, err := converter.Generate(reader, gojson.ParseJson, name, []string{"json"}, true)
	if err != nil {
		return "", err
	}
	return strings.Replace(string(byts), "[]", "", 1), nil
}
