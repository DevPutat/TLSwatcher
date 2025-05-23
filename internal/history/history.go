package history

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	"github.com/DevPutat/TLSwatcher/internal/logs"
	"github.com/DevPutat/TLSwatcher/internal/types"
)

func Read(historyPath string) types.History {

	if _, err := os.Stat(historyPath); errors.Is(err, os.ErrNotExist) {

		logs.Add(types.ErrorLog{
			Package: "history",
			Err:     err,
		})
		return types.History{}
	}

	file, err := os.Open(historyPath)
	if err != nil {
		logs.Add(types.ErrorLog{
			Package: "history",
			Err:     err,
		})
		return types.History{}
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)

	var h types.History

	json.Unmarshal(byteValue, &h)
	return h
}

func Write(historyPath string, domains []types.Domain) {

	if _, err := os.Stat(historyPath); errors.Is(err, os.ErrNotExist) {
		logs.Add(types.ErrorLog{
			Package: "history",
			Err:     err,
		})
	}
	now := time.Now()
	h := types.History{
		Domains:  domains,
		Datetime: now,
	}
	jsonString, _ := json.Marshal(h)
	os.WriteFile(historyPath, jsonString, 0644)
}
