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
	// FIXME(mdr): Currently, if you try to create the same movie twice, a new
	// entry won't be created in DynamoDB. I don't think that is in line with
	// general REST principals, since POST is neither safe nore idempotent.
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
	// FIXME(mdr): Should return the Location header as well.
	return resp.Created()
}

func main() {
	lambda.Start(Handler)
}
