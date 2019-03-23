package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the GET by year requests.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	year := req.PathParameters["year"]
	fmt.Printf("Path vars: %s\n", year)
	items, err := dao.ListByYear(year)
	if err != nil {
		// TODO(mdr): Do we really want to panic in this situation?
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}

	// Make sure the Item isn't empty
	if len(items) == 0 {
		fmt.Printf("Could not find movies with year %s\n", year)
		return resp.InternalError(req.Body)
	}

	// Log and return result
	// TODO(mdr): Refactor marshaling?
	stringItems := "["
	for i := 0; i < len(items); i++ {
		jsonItem, _ := json.Marshal(items[i])
		stringItems += string(jsonItem)
		if i != len(items)-1 {
			stringItems += ",\n"
		}
	}
	stringItems += "]\n"
	fmt.Println("Found items: ", stringItems)
	return resp.OK(stringItems)
}

func main() {
	lambda.Start(Handler)
}
