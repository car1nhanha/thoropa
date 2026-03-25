package model

type Link struct {
	Id        string `dynamodbav:"id"`
	Ip        string `dynamodbav:"ip"`
	CreatedAt int64  `dynamodbav:"created_at"`
	Accesses  int32  `dynamodbav:"accesses"`
	Original  string `dynamodbav:"original"`
}
