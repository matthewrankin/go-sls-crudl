package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/parse"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the DELETE requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	fmt.Println("Path vars: ", request.PathParameters["year"], " ", parse.Unslugify(request.PathParameters["title"]))
	err := dao.Delete(request.PathParameters["year"], parse.Unslugify(request.PathParameters["title"]))
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	return resp.Success()
}

func main() {
	lambda.Start(Handler)
}
