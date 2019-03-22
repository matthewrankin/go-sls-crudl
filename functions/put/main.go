package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
)

// Handler handles the PUT requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Log body and pass to the DAO
	fmt.Println("Received body: ", request.Body)
	item, err := dao.Put(request.Body)
	if err != nil {
		fmt.Println("Got error calling put")
		fmt.Println(err.Error())
		return events.APIGatewayProxyResponse{
			Body:       "Error",
			StatusCode: http.StatusInternalServerError,
		}, nil
	}

	// Log and return result
	fmt.Println("Updated item:  ", item)
	return events.APIGatewayProxyResponse{
		Body:       "Success\n",
		StatusCode: http.StatusOK,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
