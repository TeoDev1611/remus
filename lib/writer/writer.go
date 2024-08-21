package writer

import (
	"bytes"
	"encoding/json"
	"os"

	"github.com/TeoDev1611/remus/errors"
	toml "github.com/pelletier/go-toml/v2"
)

func WriteJsonDB(name, spaces string, empty bool, data ...map[string]interface{}) {
	if spaces == "" {
		spaces = "  "
	}

	if name == "" {
		name = "./db.json"
	}
	if empty {
		data = []map[string]interface{}{{
			"database":    "Remus",
			"awesome":     true,
			"description": "Experimental Database",
		}}
	}
	file, err := json.MarshalIndent(data, "", spaces)
	errors.CheckErrors(err)
	err2 := os.WriteFile(name, file, os.ModePerm)
	errors.CheckErrors(err2)
}

func WriteTomlDB(name string, empty bool, data ...map[string]interface{}) {
	if name == "" {
		name = "db.toml"
	}
	if empty {
		data = []map[string]interface{}{{
			"database":    "Remus",
			"awesome":     true,
			"description": "Experimental DataBase",
		}}
	}
	buf := bytes.Buffer{}
	enc := toml.NewEncoder(&buf)
	enc.SetIndentTables(true)
	enc.Encode(data)
	err2 := os.WriteFile(name, buf.Bytes(), os.ModePerm)
	errors.CheckErrors(err2)
}
