package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func main() {

	type MovieKey struct {
		Year  int    `json:"year"`
		Title string `json:"title"`
	}

	type MovieRatingCondition struct {
		Rating int `json:":val"`
	}

	config := &aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000"),
	}

	sess := session.Must(session.NewSession(config))

	svc := dynamodb.New(sess)

	key, err := dynamodbattribute.MarshalMap(MovieKey{
		Year:  2015,
		Title: "The Big New Movie",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	condition, err := dynamodbattribute.MarshalMap(MovieRatingCondition{
		Rating: 5.0,
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	input := &dynamodb.DeleteItemInput{
		Key:                       key,
		TableName:                 aws.String("Movies"),
		ConditionExpression:       aws.String("info.rating <= :val"),
		ExpressionAttributeValues: condition,
	}

	_, err = svc.DeleteItem(input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("The item has been deleted")
}
