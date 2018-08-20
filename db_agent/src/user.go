package src

import (
	"microservice_learning/protobuf/dbagent"
	"fmt"
	"context"
)


func (d *DbServer) GetOneTestUser(ctx context.Context, in *dbagent.StringValue) (*dbagent.TestUser, error) {
	fmt.Println(in.Value)
	return &dbagent.TestUser{Id:1,Account:"2",Password:"3"}, nil
}

func (d *DbServer) GetAllTestUser(ctx context.Context, in *dbagent.StringValue) (*dbagent.ListUser, error) {
	return nil, nil
}