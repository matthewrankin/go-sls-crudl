package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the GET by year requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	fmt.Println("Path vars: ", request.PathParameters["year"])
	items, err := dao.ListByYear(request.PathParameters["year"])
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}

	// Make sure the Item isn't empty
	if len(items) == 0 {
		fmt.Println("Could not find movies with year ", request.PathParameters["year"])
		return resp.InternalError(request.Body)
	}

	// Log and return result
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
