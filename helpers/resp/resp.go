package resp

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Created returns an API GatewayProxyResponse with an empty body and a Created
// HTTP status code.
func Created(location string) (events.APIGatewayProxyResponse, error) {
	hdrs := map[string]string{
		"Content-Type": "application/vnd.go-sls-crudl-movies+json",
		"Location":     location,
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Headers:    hdrs,
		Body:       "",
	}, nil
}

// NoContent returns an API GatewayProxyResponse with an empty body and a No
// Content HTTP status code.
func NoContent() (events.APIGatewayProxyResponse, error) {
	return Response("", http.StatusNoContent)
}

// OK returns an API GatewayProxyResponse with the given body and an OK HTTP
// status code.
func OK(body string) (events.APIGatewayProxyResponse, error) {
	return Response(body, http.StatusOK)
}

// InternalError returns an API GatewayProxyResponse with the given body and
// error.
func InternalError(body string) (events.APIGatewayProxyResponse, error) {
	return Response(body, http.StatusInternalServerError)
}

// Response returns an API GatewayProxyResponse with the given body and the
// given HTTP status code.
func Response(body string, code int) (events.APIGatewayProxyResponse, error) {
	hdrs := map[string]string{
		"Content-Type": "application/vnd.go-sls-crudl-movies+json",
	}
	return events.APIGatewayProxyResponse{
		StatusCode: code,
		Headers:    hdrs,
		Body:       body,
	}, nil
}
