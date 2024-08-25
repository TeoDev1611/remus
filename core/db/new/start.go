package new

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/TeoDev1611/remus/lib/files/parser"
	"github.com/TeoDev1611/remus/lib/files/writer"
	"github.com/TeoDev1611/remus/logger"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
)

var dataFiles []map[string]interface{}

func StartDB(fileInfo UIData) {
	dbFormat := promptui.Select{
		Label: "Choose the Database Format for visualization!",
		Items: []string{"json", "toml", "yaml"},
	}
	_, format, err := dbFormat.Run()
	logger.CheckErrors(err)

	defaultData := promptui.Prompt{
		Label:     "Do you want use the default data?",
		IsConfirm: true,
	}
	useDefaultData, _ := defaultData.Run()

	if strings.ToLower(useDefaultData) == "y" {
		data := parser.ParseMapToFormat(format, writer.DefaultData)
		fmt.Println(data)
		os.Exit(0)
	}

	promptFilesEditor := promptui.Prompt{
		Label: "What files do you want use?",
	}
	filesEditor, err := promptFilesEditor.Run()
	logger.CheckErrors(err)

	if strings.Contains(filesEditor, " ") {
		files := strings.Split(filesEditor, " ")
		for _, v := range files {
			auxData := []map[string]interface{}{}
			data, err := os.ReadFile(v)
			logger.CheckErrors(err)
			auxData = append(auxData, parser.ParseTextToMap(format, string(data)))
			dataFiles = append(dataFiles, auxData...)
		}
		joinedData := parser.ParseMapToFormat(format, dataFiles)
		fmt.Println(joinedData)
		switch format {
		case "json":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteJsonDB(dbPath, joinedData)
		case "toml":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteTomlDB(dbPath, joinedData)
		case "yaml":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteYamlDB(dbPath, joinedData)
		}
	} else {
		data, err := os.ReadFile(filesEditor)
		logger.CheckErrors(err)
		parsedData := parser.ParseTextToMap(format, string(data))
		fmt.Println(parsedData)

		switch format {
		case "json":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteJsonDB(dbPath, parsedData)
		case "toml":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteTomlDB(dbPath, parsedData)
		case "yaml":
			dbPath := path.Join(viper.GetString("db_dir"), fileInfo.Name)
			logger.Info(fmt.Sprintf("The database path and name is: %v", dbPath))
			writer.WriteYamlDB(dbPath, parsedData)

		}
	}
}
