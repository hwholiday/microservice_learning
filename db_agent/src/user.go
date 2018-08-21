package src

import (
	"microservice_learning/protobuf/dbagent"
	"context"
	"fmt"
)

func (d *DbServer) GetOneTestUser(ctx context.Context, req *dbagent.StringValue, rsp *dbagent.TestUser) error {
	fmt.Println(req.Value)
	return nil
}
func (d *DbServer) GetAllTestUser(ctx context.Context, req *dbagent.StringValue, rsp *dbagent.ListUser) error {
	return nil
}
