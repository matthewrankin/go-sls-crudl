package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the PUT requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// FIXME(mdr): Currently, the call to the PUT endpoint, doesn't require the
	// ID in the URL. I don't think that falls in line with general REST
	// principles.
	// Log body and pass to the DAO
	fmt.Println("Received body: ", request.Body)
	item, err := dao.Put(request.Body)
	if err != nil {
		fmt.Println("Got error calling put")
		fmt.Println(err.Error())
		return resp.InternalError("Error")
	}

	// Log and return result
	fmt.Println("Updated item: ", item)
	return resp.NoContent()
}

func main() {
	lambda.Start(Handler)
}
