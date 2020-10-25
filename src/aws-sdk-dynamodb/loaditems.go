package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Item struct {
	Year   int
	Title  string
	Plot   string
	Rating float64
}

func getItems() []Item {
	raw, err := ioutil.ReadFile("./movie_data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var items []Item
	json.Unmarshal(raw, &items)
	return items
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := dynamodb.New(sess)

	items := getItems()

	tableName := "Vijay-Movies"

	for _, item := range items {
		av, err := dynamodbattribute.MarshalMap(item)
		if err != nil {
			fmt.Println("Got error marshalling map:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		input := &dynamodb.PutItemInput{
			Item:      av,
			TableName: aws.String(tableName),
		}

		_, err = svc.PutItem(input)
		if err != nil {
			fmt.Println("Got error calling PutItem:")
			fmt.Println(err.Error())
			os.Exit(1)
		}

		year := strconv.Itoa(item.Year)

		fmt.Println("Successfully added '" + item.Title + "' (" + year + ") to table " + tableName)
	}
}
