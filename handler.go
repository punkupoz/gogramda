package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

	"github.com/punkupoz/gogramda/resolver"
	"github.com/punkupoz/gogramda/schema"
	"github.com/punkupoz/gogramda/playground"

	graphql "github.com/graph-gophers/graphql-go"
)

var mainSchema *graphql.Schema

var (
	// ErrNameNotProvided is thrown when a name is not provided
	QueryNameNotProvided = errors.New("no query was provided in the HTTP body")
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if(request.HTTPMethod == "GET") {
		return events.APIGatewayProxyResponse{
			Headers:    map[string]string{ "Content-Type": "text/html" },
			Body:       playground.Playground("GraphQL Playground", "/dev"),
			StatusCode: 200,
		}, nil
	}
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no query is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, QueryNameNotProvided
	}

	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}

	if err := json.Unmarshal([]byte(request.Body), &params); err != nil {
		log.Print("Could not decode body", err)
	}

	response := mainSchema.Exec(context, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		log.Print("Could not decode body")
	}

	return events.APIGatewayProxyResponse{
		Body:       string(responseJSON),
		StatusCode: 200,
	}, nil
}

func init() {
	root, _ := resolver.NewRoot()
	rootSchema := schema.String()
	mainSchema = graphql.MustParseSchema(rootSchema, root)
}

func main() {
	lambda.Start(Handler)
}

