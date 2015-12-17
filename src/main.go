package main

import (
	"fmt"
	
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func main() {
	fmt.Println("Hello, world")
	
	sqs_client := sqs.New(session.New(), aws.NewConfig().WithRegion("us-west-2"))
}