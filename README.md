# Go Lambda API with SAM CLI

## Overview

This project provides a Go Lambda function deployed with the **AWS Serverless Application Model (SAM)** and **API Gateway**. It handles incoming HTTP requests via **API Gateway** and processes the request with the Lambda function written in **Go**. The Lambda function parses query parameters and returns a message.

## Prerequisites

Before you begin, ensure you have the following tools installed:

1. **AWS CLI**: [Installation Guide](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2.html)
2. **AWS SAM CLI**: [Installation Guide](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)
3. **Docker**: (Required for local testing and building Lambda containers)
   - [Docker Installation Guide](https://docs.docker.com/get-docker/)
4. **GoLang**: (For building the Lambda function locally)
   - [GoLang Installation Guide](https://go.dev/dl/)

5. **AWS Account**: You need an active AWS account to deploy the Lambda function and API Gateway.

## Project Setup

### Step 1: Clone the Repository

Clone the repository to your local machine:

```bash
git clone https://github.com/your-repo/GoLambdaApp.git
cd GoLambdaApp
```

### Step 2: Install Dependencies

Navigate to the `hello-world` directory where the Lambda code is stored and run:

```bash
go mod init GoLambdaApp
go get github.com/aws/aws-lambda-go/lambda
```

### Step 3: Build the Go Application

To build the Go Lambda function, use the following command:

```bash
$env:GOOS="linux"; $env:GOARCH="amd64"; go build -o bootstrap main.go

sam build
```

This command will package the Go Lambda function and prepare it for deployment.

---

## Deployment with AWS SAM CLI

### Step 1: Deploy the Application

You can deploy the Lambda function and API Gateway by using the `sam deploy --guided` command.

```bash
sam deploy --guided
```

#### During deployment, you will be prompted for:
- **Stack name**: Name of your CloudFormation stack (e.g., `GoLambdaApp`).
- **AWS Region**: The AWS region for the deployment (e.g., `us-west-2`).
- **Confirm changes before deploy**: Yes or No, depending on your preference.
- **Allow SAM CLI to create IAM roles with the required permissions**: Yes.

Once the deployment is successful, you will receive an **API Gateway URL** that you can use to invoke your Lambda function.

---

## Accessing the Lambda API Endpoint

Once deployed, the Lambda function is exposed via an API Gateway endpoint. You can test the endpoint using `curl`, a browser, or Postman.

### API Endpoint

The Lambda function is available at the following URL format:

```bash
https://{api-id}.execute-api.{region}.amazonaws.com/prod/hello
```

- Replace `{api-id}` with the API Gateway ID provided after deployment.
- Replace `{region}` with your AWS region (e.g., `us-west-2`).

### Example `curl` Command

To send a GET request with a query parameter (`name`), use the following `curl` command:

```bash
curl "https://{api-id}.execute-api.{region}.amazonaws.com/prod/hello?name=John"
```

**Response:**

```text
Hello, John!
```

If no `name` parameter is provided, the default value `World` will be used:

```bash
curl "https://{api-id}.execute-api.{region}.amazonaws.com/prod/hello"
```

**Response:**

```text
Hello, World!
```

---

## Lambda Function Code Explanation

The Lambda function is implemented in **Go** (`main.go`) and is designed to:

1. Parse query parameters (`name`) from the incoming HTTP request.
2. If `name` is provided, return `"Hello, {name}!"`.
3. If `name` is not provided, return `"Hello, World!"`.

### Lambda Handler (`main.go`)

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    log.Println("Request received:", request)

    name := request.QueryStringParameters["name"]
    if name == "" {
        name = "World"
    }

    message := fmt.Sprintf("Hello, %s!", name)

    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Body:       message,
    }, nil
}

func main() {
    lambda.Start(handler)
}
```

### SAM Template (`template.yaml`)

The **SAM template** defines the Lambda function and API Gateway configuration:

```yaml
AWSTemplateFormatVersion: '2010-09-09'
Transform: AWS::Serverless-2016-10-31
Description: Go Lambda function with API Gateway

Resources:
  GoLambdaApi:
    Type: AWS::Serverless::Api
    Properties:
      Name: GoLambdaApi
      StageName: prod
      Cors:
        AllowOrigins: "*"
        AllowMethods: "GET,POST"

  GoLambdaFunction:
    Type: AWS::Serverless::Function
    Properties:
      FunctionName: GoLambdaFunction
      Handler: main
      Runtime: provided.al2
      CodeUri: ./hello-world
      MemorySize: 128
      Timeout: 10
      Events:
        ProxyApi:
          Type: Api
          Properties:
            Path: /hello
            Method: get
            RestApiId: !Ref GoLambdaApi

Outputs:
  GoLambdaFunctionOutput:
    Description: "Lambda Function ARN"
    Value: !GetAtt GoLambdaFunction.Arn
  GoApiUrl:
    Description: "API Gateway URL"
    Value: !Sub "https://${GoLambdaApi}.execute-api.${AWS::Region}.amazonaws.com/prod/hello"
```

---

## Cleanup

To remove the deployed resources from AWS, use the following command:

```bash
sam delete
```

This will delete the CloudFormation stack and all associated resources like Lambda and API Gateway.

---

## Troubleshooting

- **CORS Issues**: If you face CORS issues, ensure your API Gateway has CORS enabled in the SAM template.
- **Lambda Timeout**: If the Lambda function times out, increase the `Timeout` value in the `template.yaml`.
- **AWS Credentials**: Ensure your AWS CLI is configured with the correct credentials by running `aws configure`.

---

## Contributing

Feel free to fork this repository, make improvements, and create pull requests. Any contributions are appreciated!

---

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

