package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/parse"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the DELETE requests.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	year := req.PathParameters["year"]
	title := parse.Unslugify(req.PathParameters["title"])
	fmt.Printf("Path vars: %s - %s\n", year, title)
	err := dao.Delete(year, title)
	if err != nil {
		// TODO(mdr): Do we really want to panic in this situation?
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}
	return resp.NoContent()
}

func main() {
	lambda.Start(Handler)
}
