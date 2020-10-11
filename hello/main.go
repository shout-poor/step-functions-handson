package main

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shout-poor/step-functions-handson/domain"
)

func Handler(ctx context.Context) (resp domain.DataTransferObject, error) {
	return domain.DataTransferObject{ID: 1, Name: "John doe"}, nil
}

func main() {
	lambda.Start(Handler)
}
