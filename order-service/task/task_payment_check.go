package task

import (
	"encoding/base64"
	"encoding/json"

	"github.com/labstack/gommon/log"
	"github.com/skyapps-id/edot-test/order-service/wrapper/shop_warehouse_service"
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

func SendWebhook(b64payload string) (bool, error) {
	payload := shop_warehouse_service.ProductStockReductionRequest{}
	DecodeToTask(b64payload, &payload)

	log.Info("=========> Task Run", payload)

	return false, nil
}
