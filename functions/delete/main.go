package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
)

// Parse slug into a space separated string
func parseSlug(orig string) (retval string) {
	retval = strings.Replace(orig, "-", " ", -1)
	retval = strings.Replace(retval, "+", " ", -1)
	return retval
}

// Handler handles the DELETE requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	fmt.Println("Path vars: ", request.PathParameters["year"], " ", parseSlug(request.PathParameters["title"]))
	err := dao.Delete(request.PathParameters["year"], parseSlug(request.PathParameters["title"]))
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	return events.APIGatewayProxyResponse{
		Body:       "Success\n",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
