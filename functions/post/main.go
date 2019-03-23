package main

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/matthewrankin/go-sls-crudl/helpers/dao"
	"github.com/matthewrankin/go-sls-crudl/helpers/parse"
	"github.com/matthewrankin/go-sls-crudl/helpers/resp"
)

// Handler handles the POST request.
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// FIXME(mdr): Currently, if you try to create the same movie twice, a new
	// entry won't be created in DynamoDB. I don't think that is in line with
	// general REST principals, since POST is neither safe nore idempotent.
	// Log body and pass to the DAO
	fmt.Println("Received body: ", req.Body)
	item, err := dao.Post(req.Body)
	if err != nil {
		fmt.Println("Got error calling post")
		fmt.Println(err.Error())
		return resp.InternalError("Error")
	}

	// Log and return result
	fmt.Println("Wrote item: ", item)
	log.Printf("req = %#v", req)
	baseURL := req.Headers["Host"]
	stage := req.RequestContext.Stage
	title := parse.Slugify(item.Title)
	location := fmt.Sprintf("https://%s/%s%s/%d/%s", baseURL, stage, req.Path, item.Year, title)
	return resp.Created(location)
}

func main() {
	lambda.Start(Handler)
}
