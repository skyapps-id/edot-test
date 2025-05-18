package task

import (
	"encoding/base64"
	"encoding/json"
)

func DecodeToTask(msg string, task interface{}) (err error) {
	decodedstg, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return
	}
	msgBytes := decodedstg
	err = json.Unmarshal(msgBytes, task)
	if err != nil {
		return
	}
	return
}
