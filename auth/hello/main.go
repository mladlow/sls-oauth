package main

import (
  "bytes"
  "context"
  "encoding/json"

  "github.com/aws/aws-sdk-go/aws"

  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"

  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
)

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
  var buf bytes.Buffer

  svc := dynamodb.New(session.New())
  input := &dynamodb.PutItemInput{
    Item: map[string]*dynamodb.AttributeValue{
      "user_id": {
        S: aws.String("abc123"),
      },
    },
    ReturnConsumedCapacity: aws.String("TOTAL"),
    TableName: aws.String("sls-oauth"),
  }
  _, err := svc.PutItem(input)
  if err != nil {
    return Response{StatusCode: 500}, err
  }

  body, err := json.Marshal(map[string]interface{}{
    "message": "Go Serverless v1.0! Your function inserted an item!",
  })
  if err != nil {
    return Response{StatusCode: 404}, err
  }
  json.HTMLEscape(&buf, body)

  resp := Response{
    StatusCode:      200,
    IsBase64Encoded: false,
    Body:            buf.String(),
    Headers: map[string]string{
      "Content-Type":           "application/json",
      "X-MyCompany-Func-Reply": "hello-handler",
    },
  }

  return resp, nil
}

func main() {
  lambda.Start(Handler)
}
