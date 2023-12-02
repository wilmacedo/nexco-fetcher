package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/wilmacedo/nexco-fetcher/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
