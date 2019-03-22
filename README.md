# go-sls-crudl

This project riffs off of the [Dynamo DB Golang samples][1] and the
[Serverless Framework Go example][2] to create an example of how to
build a simple API Gateway -> Lambda -> DynamoDB set of methods.

## Code Organization

Note that, instead of using the `create_table.go` to set up the initial
table like the AWS code example does, the resource building mechanism
that Serverless provides is used.  Individual code is organized as
follows:

- `functions/get/main.go` - GET method for reading a specific item
- `functions/list-by-year/main.go` - GET method for listing all or a
  subset of items
- `functions/post/main.go` - POST method for creating a new item
- `functions/put/main.go` - PUT method for updating an existing item
- `functions/delete/main.go` - DELETE method for deleting a specific
  item
- `img/*` - Images of DynamoDB tables to make this Readme easier to
  follow
- `helpers/dao/moviedao.go` - DAO wrapper around the DynamoDB calls
- `helpers/parse/parse.go` - Helper functions related to parsing, such
  as Unslugify to convert a slug into a space separated string.
- `helpers/resp/resp.go` - Helper functions to simplify providing an
  APIGatewayProxyResponse.
- `data/XXX.json` - Set of sample data files for POST and PUT actions
- `Makefile` - Used for testing, building, and deploying the functions.
- `serverless.yml` - Defines the initial table, function defs, and API
  Gateway events

Compared to [nerdguru/go-sls-crudl][3] layout, I've moved each `main`
package into its own folder, so that the `main()` func isn't being
redeclared in multiple files in the same directory.

## Set Up

If you are a Serverless Framework rookie, [follow the installation
instructions here][sls-install].  If you are a grizzled vet, be sure
that you have v1.26 or later as that's the version that introduces
Golang support.  You'll also need to [install Go][go-install].

When both of those tasks are done, and after you've cloned the
repository, you can run `make` to see the available commands to build,
check, clean, and deploy.

When done deploying using `make deploy`, you can find the new DynamoDB
table in the AWS Console, which should initially look like this:

![Initial DynamoDB Table](/img/initialDynamoDBTable.jpg)

and your `<base URL>` will be of the format
'https://XXXXXXXXXX.execute-api.us-east-1.amazonaws.com/dev/movies'
where `XXXXXXXXXX` will be some random string generated by AWS.

The development cycle would then be:

- Make changes to the .go files
- Run `make deploy` to compile and deploy the binaries
- Use `curl` to then interrogate the API as described below

## Using

Once deployed and substituting your `<base URL>` the following CURL
commands can be used to interact with the resulting API, whose results
can be confirmed in the DynamoDB console

### POST

```bash
curl -X POST https:<base URL> -d @data/post1.json
```
Which should result in the DynamoDB table looking like this:

![First Post DynamoDB Table](/img/firstPostDynamoDBTable.jpg)

Rinse/repeat for other data files to yeild:

![All Posts DynamoDB Table](/img/allPostsDynamoDBTable.jpg)

### GET Specific Item

Using the year and title (replacing spaces wiht '-' or '+'), you can now
obtain an item as follows (prettified output):

```bash
curl https://<base URL>/2013/Hunger-Games-Catching-Fire
{
  "year": 2013,
  "title": "Hunger Games Catching Fire",
  "info": {
    "plot": "Katniss Everdeen and Peeta Mellark become targets of the
             Capitol after their victory in the 74th Hunger Games sparks
             a rebellion in the Districts of Panem.",
    "rating": 7.6
  }
}
```

### GET a List of Items

You can list items by year as follows (prettified output):

```bash
curl https://<base URL>/2013
[
  {
    "year": 2013,
    "title": "Hunger Games Catching Fire",
    "info": {
      "plot": "",
      "rating": 0
    }
  },
  {
    "year": 2013,
    "title": "Turn It Down Or Else",
    "info": {
      "plot": "",
      "rating": 0
    }
  }
]
```

### DELETE Specific Item

Using the same year and title specifiers, you can delete as follows:

```bash
curl -X DELETE https://<base URL>/2013/Hunger-Games-Catching-Fire
```

Which should result in the DynamoDB table looking like this:

![First Delete DynamoDB Table](/img/firstDeleteDynamoDBTable.jpg)

### UPDATE Specific Item

You can update as follows:

```bash
curl -X PUT https:<base URL> -d @data/put3.json
```

Which should result in the DynamoDB table looking like this:

![First Update DynamoDB Table](/img/firstUpdateDynamoDBTable.jpg)


[1]: https://github.com/awsdocs/aws-doc-sdk-examples/tree/master/go/example_code/dynamodb
[2]: https://serverless.com/blog/framework-example-golang-lambda-support/
[3]: https://github.com/nerdguru/go-sls-crudl
[sls-install]: https://serverless.com/blog/anatomy-of-a-serverless-app/#setup
[go-install]: https://golang.org/doc/install