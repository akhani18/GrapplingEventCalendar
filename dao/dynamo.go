package dao

import (
	"fmt"
	"os"

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

// initialize dynamodb client

func GetUpcomingCompetitionsInState() {
	params := &dynamodb.ScanInput{
		TableName: aws.String("competitions"),
	}

	result, err := db.Scan(params)
	if err != nil {
		exitWithError(fmt.Errorf("failed to make scan API call, %v", err))
	}

	comps := []Competition{}

	// Unmarshal the Items field in the result value to the Item Go type.
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &comps)
	if err != nil {
		exitWithError(fmt.Errorf("failed to unmarshal scan result items, %v", err))
	}

	// Print out the items returned
	for i, comp := range comps {
		fmt.Printf("%d. name: %s, date: %s, city: %s, state: %s", i, comp.Name, comp.Date, comp.City, comp.State)
	}
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
