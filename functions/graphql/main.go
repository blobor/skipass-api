package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/blobor/skipass-api/functions/graphql/resolver"
	"github.com/blobor/skipass-api/functions/graphql/schema"
	"github.com/graph-gophers/graphql-go"
)

var mainSchema *graphql.Schema

type queryParams struct {
	Query         string                 `json:"query"`
	OperationName string                 `json:"operationName"`
	Variables     map[string]interface{} `json:"variables"`
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	if len(request.Body) < 1 {
		log.Print("Request have empty body")

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
		}, nil
	}

	var params queryParams
	err := json.Unmarshal([]byte(request.Body), &params)

	if err != nil {
		log.Printf("Could not decode body, %s", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	res := mainSchema.Exec(ctx, params.Query, params.OperationName, params.Variables)

	jsonRes, err := json.Marshal(res)

	if err != nil {
		log.Printf("Could not encode body, %s", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Headers: map[string]string{
				"Access-Control-Allow-Origin": "*",
			},
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       string(jsonRes),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
	}, nil
}

func init() {
	mainSchema = graphql.MustParseSchema(schema.String(), resolver.NewRoot())
}

func main() {
	lambda.Start(Handler)
}
