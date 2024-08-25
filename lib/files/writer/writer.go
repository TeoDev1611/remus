package writer

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/TeoDev1611/remus/logger"
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

func WriteJsonDB(name string, data interface{}) {
	file, err := json.MarshalIndent(data, "", " ")
	logger.CheckErrors(err)
	err2 := afs.WriteFile(name, file, os.ModePerm)
	logger.CheckErrors(err2)
}

func WriteTomlDB(name string, data interface{}) {
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	enc.Encode(data)
	err2 := afs.WriteFile(name, buf.Bytes(), os.ModePerm)
	logger.CheckErrors(err2)
}

func WriteYamlDB(name string, data interface{}) {
	file, err := yaml.Marshal(data)
	logger.CheckErrors(err)
	err2 := afs.WriteFile(name, file, os.ModePerm)
	logger.CheckErrors(err2)
}
