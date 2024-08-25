package data

import (
	"encoding/json"
	"fmt"

	"github.com/TeoDev1611/remus/logger"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

// Afero Utils
var (
	fs  = afero.NewMemMapFs()
	afs = &afero.Afero{Fs: fs}
)

func find(obj interface{}, key string) (interface{}, bool) {
	mobj, ok := obj.(map[string]interface{})
	if !ok {
		return nil, false
	}

	for k, v := range mobj {
		if k == key {
			return v, true
		}

		if m, ok := v.(map[string]interface{}); ok {
			if res, ok := find(m, key); ok {
				return res, true
			}
		}
		if va, ok := v.([]interface{}); ok {
			for _, a := range va {
				if res, ok := find(a, key); ok {
					return res, true
				}
			}
		}
	}

	return nil, false
}

func GetAllDataFileJson(filename string) []interface{} {
	file, err := afs.ReadFile(filename)
	logger.CheckErrors(err)
	var data []interface{}
	err2 := json.Unmarshal(file, &data)
	logger.CheckErrors(err2)
	return data
}

func GetSingleKeyJson(filename, key string) string {
	data := GetAllDataFileJson(filename)
	val, info := find(data, key)
	if info {
		return fmt.Sprintf("%v", val)
	}
	return "No data"
}

func GetAllDataFileToml(filename string) []interface{} {
	file, err := afs.ReadFile(filename)
	logger.CheckErrors(err)
	var data []interface{}
	err2 := toml.Unmarshal(file, &data)
	logger.CheckErrors(err2)
	return data
}

func GetSingleKeyToml(filename, key string) string {
	data := GetAllDataFileToml(filename)
	val, info := find(data, key)
	if info {
		return fmt.Sprintf("%v", val)
	}
	return "No data"
}

func GetAllDataFileYaml(filename string) []interface{} {
	file, err := afs.ReadFile(filename)
	logger.CheckErrors(err)
	var data []interface{}
	err2 := yaml.Unmarshal(file, &data)
	logger.CheckErrors(err2)
	return data
}

func GetSingleKeyYaml(filename, key string) string {
	data := GetAllDataFileYaml(filename)
	val, info := find(data, key)
	if info {
		return fmt.Sprintf("%v", val)
	}
	return "No data"
}
