package repository

import (
	"context"
	"thoropa/internal/model"

	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type LinkRepository struct {
	client *dynamodb.Client
	table  string
}

func NewLinkRepository(client *dynamodb.Client) *LinkRepository {
	return &LinkRepository{
		client: client,
		table:  "links",
	}
}

func (r *LinkRepository) Create(ctx context.Context, link *model.Link) error {
	item, err := attributevalue.MarshalMap(link)

	if err != nil {
		return err
	}

	_, err = r.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &r.table,
		Item:      item,
	})

	return err
}
