package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/punkupoz/gogramda"
)

func main() {
	lambda.Start(gogramda.Handler)
}
