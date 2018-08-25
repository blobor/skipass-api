package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHandler(t *testing.T) {
	tests := []struct {
		request events.APIGatewayProxyRequest
		expect  events.APIGatewayProxyResponse
		context context.Context
		error   error
	}{
		{
			request: events.APIGatewayProxyRequest{
				Body: "",
			},
			expect: events.APIGatewayProxyResponse{
				StatusCode: 400,
			},
			context: context.TODO(),
		},
	}

	for _, test := range tests {
		response, err := Handler(test.context, test.request)
		assert.IsType(t, test.error, err)
		assert.Equal(t, test.expect, response)
	}
}
