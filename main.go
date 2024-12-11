package main

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Parse form data from the request body
	parsedFormData, err := url.ParseQuery(request.Body)
	if err != nil {
		log.Printf("Error parsing form data: %s", err)
		return events.APIGatewayProxyResponse{StatusCode: 400}, err
	}

	// Access POST parameters
	id := parsedFormData.Get("Id")
	name := parsedFormData.Get("Name")

	log.Printf("Received Id: %s", id)
	log.Printf("Received Name: %s", name)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       fmt.Sprintf(`Id: %s, Name: %s`, id, name),
	}, nil
}

func main() {
	lambda.Start(handler)
}
