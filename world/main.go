package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/shout-poor/step-functions-handson/domain"
)

func Handler(ctx context.Context, event domain.DataTransferObject) error {
	fmt.Sprintf("%+v", event)
	return nil
}

func main() {
	lambda.Start(Handler)
}
