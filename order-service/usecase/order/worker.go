package order

import (
	"context"
	"encoding/base64"
	"time"

	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/skyapps-id/edot-test/order-service/pkg/logger"
	"github.com/skyapps-id/edot-test/order-service/pkg/tracer"
)

func (uc *usecase) SendTask(ctx context.Context, reqJSON []byte) {
	_, span := tracer.Define().Start(ctx, "OrderUsecaseWorker.SendTask")
	defer span.End()

	b64EncodedReq := base64.StdEncoding.EncodeToString(reqJSON)
	eta := time.Now().UTC().Add(time.Second * 5)
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

	_, err := uc.worker.SendTask(&task)
	if err != nil {
		logger.Log.Error(err.Error())
	}
}
