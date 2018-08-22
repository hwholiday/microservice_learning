package src

import (
	"context"
	"microservice_learning/protobuf/logagent"
	"fmt"
)

func (l *LogServer)StartLog(ctx context.Context, log *logagent.Log) error {
	fmt.Printf("Got event %+v\n", log)
	return nil
}
