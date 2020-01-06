package dao

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	region = os.Getenv("AWS_REGION")
	//table  = os.Getenv("COMPETITONS_TABLE")
	db = dynamodb.New(session.New(), aws.NewConfig().WithRegion(region))
)

/*
aws dynamodb query --table-name competitions \
--index-name state-eventDate-index \
--key-condition-expression "#st = :st AND #ed > :ed" \
--expression-attribute-names '{ "#st": "state", "#ed": "eventDate" }' \
--expression-attribute-values '{ ":st": {"S": "oregon"}, ":ed": {"S": "2020-04-05"} }' \
--region us-west-2 \
--profile ankit-dev
*/
// make state lowercase
//
// TO DO: pagination?
func GetUpcomingCompetitionsInState(state string) []*Competition {
	state = strings.ToLower(state)
	currentDate := time.Now().Format("2006-01-02")

	params := &dynamodb.QueryInput{
		KeyConditionExpression: aws.String("#st = :st AND #ed > :ed"),
		ExpressionAttributeNames: map[string]*string{
			"#st": aws.String("state"),
			"#ed": aws.String("eventDate"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":st": {
				S: aws.String(state),
			},
			":ed": {
				S: aws.String(currentDate),
			},
		},
		TableName: aws.String("competitions"),          // TO DO: Get from env var
		IndexName: aws.String("state-eventDate-index"), // TO DO: Get from env var
	}

	result, err := db.Query(params)
	if err != nil {
		exitWithError(fmt.Errorf("failed to make scan API call, %v", err))
	}

	comps := []*Competition{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &comps)
	if err != nil {
		exitWithError(fmt.Errorf("failed to unmarshal scan result items, %v", err))
	}

	// Print out the items returned
	for i, comp := range comps {
		fmt.Printf("%d. name: %s, date: %s, city: %s, state: %s", i, comp.Name, comp.Date, comp.City, comp.State)
	}

	return comps
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
