package gogramda

import (
	"context"
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"github.com/punkupoz/gogramda/playground"
	"github.com/punkupoz/gogramda/resolver"
	"github.com/punkupoz/gogramda/schema"

	graphql "github.com/graph-gophers/graphql-go"
)

var mainSchema *graphql.Schema

var (
	// ErrNameNotProvided is thrown when a name is not provided
	QueryNameNotProvided = errors.New("no query was provided in the HTTP body")
	IncorrectHttpMethod  = errors.New("accepts only GET and POST requests")
)

func Handler(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if request.HTTPMethod == "GET" {
		return events.APIGatewayProxyResponse{
			Headers: map[string]string{"Content-Type": "text/html"},
			Body:    playground.Playground("GraphQL Playground", "/dev"),

			StatusCode: 200,
		}, nil
	}

	if request.HTTPMethod == "POST" {
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

	return events.APIGatewayProxyResponse{}, IncorrectHttpMethod
}

func init() {
	root, _ := resolver.NewRoot()
	rootSchema := schema.String("./schema")
	mainSchema = graphql.MustParseSchema(rootSchema, root)
}
