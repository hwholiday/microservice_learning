package src

import (
	"microservice_learning/protobuf/dbagent"
	"context"
	"fmt"
	"time"
)

func (d *DbServer) GetOneTestUser(ctx context.Context, req *dbagent.StringValue, rsp *dbagent.TestUser) error {
	rsp.Id=1
	rsp.Account=fmt.Sprint(time.Now().Format("2006-01-02 15:04:05"))
	rsp.Password="123"
	fmt.Println(rsp.Account)
	return nil
}
func (d *DbServer) GetAllTestUser(ctx context.Context, req *dbagent.StringValue, rsp *dbagent.ListUser) error {
	return nil
}
