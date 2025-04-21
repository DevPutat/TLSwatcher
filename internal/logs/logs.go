package logs

import (
	"errors"
	"os"

	"github.com/DevPutat/TLSwatcher/internal/types"
)

func Add(errLog types.ErrorLog) {
	file, err := os.OpenFile(types.LogFilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer file.Close()
	file.WriteString(errLog.String() + "\n")
}

func CreateLogFile() {
	if _, err := os.Stat(types.LogFilePath); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(types.LogFilePath)
		if err == nil {
			file.Close()
		}
	}
}
