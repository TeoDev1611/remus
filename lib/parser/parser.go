package parser

import (
	"encoding/json"

	"github.com/TeoDev1611/remus/errors"
	toml "github.com/pelletier/go-toml/v2"
)

func ParseJsonToMap(text string) map[string]interface{} {
	byteText := []byte(text)
	var data map[string]interface{}
	err := json.Unmarshal(byteText, &data)
	errors.CheckErrors(err)
	return data
}
func ParseTomlToMap(text []byte) map[string]interface{} {
	var data map[string]interface{}
	err := toml.Unmarshal(text, &data)
	errors.CheckErrors(err)
	return data
}
