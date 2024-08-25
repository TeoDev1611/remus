package new

import (
	"fmt"
	"time"

	"github.com/TeoDev1611/remus/logger"
	"github.com/google/uuid"
	"github.com/manifoldco/promptui"
)

type UIData struct {
	Name     string
	Password string
	UID      string
	Date     string
}

func GetInfo(name, password string, prompt bool) UIData {
	if prompt && name == "" && password == "" {
		namePrompt := promptui.Prompt{
			Label:   "What's the name of the database",
			Default: "remus_example",
		}
		nameInfo, errName := namePrompt.Run()
		logger.CheckErrors(errName)

		passPrompt := promptui.Prompt{
			Label:   "What's the password for the database?",
			Default: "",
			Mask:    '*',
		}
		passwordInfo, errPass := passPrompt.Run()
		logger.CheckErrors(errPass)

		name = nameInfo
		password = passwordInfo
	}
	uid := uuid.New()
	now := time.Now()

	return UIData{
		Name:     fmt.Sprintf("%v-%v", name, uid.ID()),
		Password: password,
		UID:      uid.String(),
		Date:     now.Format(time.RFC822),
	}
}
