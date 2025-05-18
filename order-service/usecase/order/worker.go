package order

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"time"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

func (uc *usecase) OrderCheck(ctx context.Context, payload OrderCancelRequest) {
	_, span := tracer.Define().Start(ctx, "OrderUsecaseWorker.OrderCheck")
	defer span.End()

	reqJSON, err := json.Marshal(payload)
	if err != nil {
		logger.Log.Error(err.Error())
	}

	b64EncodedReq := base64.StdEncoding.EncodeToString(reqJSON)
	eta := time.Now().UTC().Add(time.Second * 10)
	task := tasks.Signature{
		Name: "send_webhook",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: b64EncodedReq,
			},
		},
		Headers: map[string]interface{}{
			"test": "test",
		},
		ETA:        &eta,
		RetryCount: 3,
	}

	_, err = uc.worker.SendTask(&task)
	if err != nil {
		logger.Log.Error(err.Error())
	}
}
