package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Parse slug into a space separated string
func parseSlug(orig string) (retval string) {
	retval = strings.Replace(orig, "-", " ", -1)
	retval = strings.Replace(retval, "+", " ", -1)
	return retval
}

// Handler handles the GET requests.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Make the call to the DAO with params found in the path
	fmt.Println("Path vars: ", request.PathParameters["year"], " ", parseSlug(request.PathParameters["title"]))
	item, err := dao.GetByYearTitle(request.PathParameters["year"], parseSlug(request.PathParameters["title"]))
	if err != nil {
		panic(fmt.Sprintf("Failed to find Item, %v", err))
	}

	// Make sure the Item isn't empty
	if item.Year <= 0 {
		fmt.Println("Could not find movie")
		return resp.InternalError(request.Body)
	}

	// Log and return result
	jsonItem, _ := json.Marshal(item)
	stringItem := string(jsonItem) + "\n"
	fmt.Println("Found item: ", stringItem)
	return resp.OK(stringItem)
}

func main() {
	lambda.Start(Handler)
}
