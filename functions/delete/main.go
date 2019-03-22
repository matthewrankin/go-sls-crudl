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
	year := request.PathParameters["year"]
	title := parse.Unslugify(request.PathParameters["title"])
	fmt.Printf("Path vars: %s - %s\n", year, title)
	err := dao.Delete(year, title)
	if err != nil {
		// TODO(mdr): Do we really want to panic in this situation?
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	return resp.Success()
}

func main() {
	lambda.Start(Handler)
}
