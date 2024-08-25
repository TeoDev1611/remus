package parser

import (
	"encoding/json"

	"github.com/TeoDev1611/remus/logger"
	"github.com/pelletier/go-toml/v2"
	"github.com/tidwall/pretty"
	"gopkg.in/yaml.v3"
)

func ParseMapToFormat(format string, data interface{}) string {
	var formatted string
	switch format {
	case "toml":
		format, err := toml.Marshal(data)
		logger.CheckErrors(err)
		formatted = string(format)
	case "json":
		format, err := json.MarshalIndent(data, "", " ")
		logger.CheckErrors(err)
		formatted = string(pretty.Pretty(format))
	case "yaml":
		format, err := yaml.Marshal(data)
		logger.CheckErrors(err)
		formatted = string(format)
	default:
		logger.NewError("No exist this format parser for the database", true)
	}
	return formatted
}

func ParseTextToMap(format, data string) map[string]interface{} {
	var parsed map[string]interface{}
	switch format {
	case "json":
		err := json.Unmarshal([]byte(data), &parsed)
		logger.CheckErrors(err)
	case "toml":
		err := toml.Unmarshal([]byte(data), &parsed)
		logger.CheckErrors(err)
	case "yaml":
		err := yaml.Unmarshal([]byte(data), &parsed)
		logger.CheckErrors(err)
	}
	return parsed
}
