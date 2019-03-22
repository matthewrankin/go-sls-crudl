package resp

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Success returns an API GatewayProxyResponse with the string "Success" in the
// body and an OK HTTP status code.
func Success() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       "Success\n",
		StatusCode: http.StatusOK,
	}, nil
}

// OK returns an API GatewayProxyResponse with the given body and an OK HTTP
// status code.
func OK(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: http.StatusOK,
	}, nil
}

// InternalError returns an API GatewayProxyResponse with the given body and
// error.
func InternalError(body string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       body,
		StatusCode: http.StatusInternalServerError,
	}, nil
}
