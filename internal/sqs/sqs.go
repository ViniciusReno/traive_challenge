package sqs

import (
	"github.com/ViniciusReno/traive/internal/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var Client *sqs.Client

func init() {
	Client = sqs.NewFromConfig(aws.AWSSession())
}
