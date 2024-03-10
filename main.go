package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/deibyssoca/ds-E-Commerce/awsgo"
	"github.com/deibyssoca/ds-E-Commerce/bd"
	"github.com/deibyssoca/ds-E-Commerce/handlers"
)

func main() {
	lambda.Start(executeLambda)
}

func executeLambda(ctx context.Context, request events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {
	awsgo.InicializeAWS()
	if !ValidateParam() {
		panic("Error in the parameters. You must send the following parameters: SecretName and UrlPrefix")
	}
	var res *events.APIGatewayProxyResponse

	// el -1 es para remplazar todo lo que encuentra
	path := strings.Replace(request.RawPath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTP.Method
	body := request.Body
	header := request.Headers

	bd.ReadSecret()

	status, message := handlers.Handlers(path, method, body, header, request)

	headerResp := map[string]string{
		"Content-Type": "application/json",
	}
	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headerResp,
	}
	return res, nil
}

// valido que me manden el parametro
func ValidateParam() bool {
	var parameterExist bool
	if _, parameterExist = os.LookupEnv("SecretName"); !parameterExist {
		return parameterExist
	}

	if _, parameterExist = os.LookupEnv("UrlPrefix"); !parameterExist {
		return parameterExist
	}
	return parameterExist
}
