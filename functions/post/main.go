package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the POST request.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Log body and pass to the DAO
	fmt.Println("Received body: ", request.Body)
	item, err := dao.Post(request.Body)
	if err != nil {
		fmt.Println("Got error calling post")
		fmt.Println(err.Error())
		return resp.InternalError("Error")
	}

	// Log and return result
	fmt.Println("Wrote item: ", item)
	return resp.Success()
}

func main() {
	lambda.Start(Handler)
}
