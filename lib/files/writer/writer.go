package writer

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/TeoDev1611/remus/errors"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

var (
	fs          = afero.NewMemMapFs()
	afs         = &afero.Afero{Fs: fs}
	DefaultData = []map[string]interface{}{{
		"database":    "Remus",
		"awesome":     true,
		"description": "Experimental Database",
	}}
)

func WriteJsonDB(name string, empty bool, data ...map[string]interface{}) {
	if name == "" {
		name = "./db.json"
	}
	if empty {
	}
	file, err := json.MarshalIndent(data, "", " ")
	errors.CheckErrors(err)
	err2 := afs.WriteFile(name, file, os.ModePerm)
	errors.CheckErrors(err2)
}

func WriteTomlDB(name string, empty bool, data ...map[string]interface{}) {
	if name == "" {
		name = "db.toml"
	}
	if empty {
		data = DefaultData
	}
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	enc.Encode(data)
	err2 := afs.WriteFile(name, buf.Bytes(), os.ModePerm)
	errors.CheckErrors(err2)
}

func WriteYamlDB(name string, empty bool, data ...map[string]interface{}) {
	if name == "" {
		name = "db.yaml"
	}
	if empty {
		data = DefaultData
	}
	file, err := yaml.Marshal(data)
	errors.CheckErrors(err)
	err2 := afs.WriteFile(name, file, os.ModePerm)
	errors.CheckErrors(err2)
}
