package database

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoClient(ctx context.Context) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("DYNAMO_LOCAL") == "true" {
		cfg.Region = "us-east-1"

		cfg.BaseEndpoint = aws.String("http://localhost:8000")

		// 👉 AQUI está o segredo
		cfg.Credentials = aws.NewCredentialsCache(
			credentials.NewStaticCredentialsProvider("dummy", "dummy", ""),
		)
	}

	return dynamodb.NewFromConfig(cfg)
}
