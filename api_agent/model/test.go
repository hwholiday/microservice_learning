package model

import (
	"microservice_learning/api_agent/server"
	"context"
	"microservice_learning/protobuf/dbagent"
)

func TestModel(data string) (*dbagent.TestUser, error) {
	return server.Client.DbAgent.GetOneTestUser(context.TODO(), &dbagent.StringValue{Value: data})
}
