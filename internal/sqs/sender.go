package sqs

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/sirupsen/logrus"
)

const delaySeconds = 10

type SQSSenderService interface {
	GetQueueURL(ctx context.Context, queueName string) (*sqs.GetQueueUrlOutput, error)
	SendMessage(ctx context.Context, queueName string, body interface{}) (*sqs.SendMessageOutput, error)
}

type service struct{}

func NewService() SQSSenderService {
	return &service{}
}

func (s service) SendMessage(ctx context.Context, queueName string, body interface{}) (*sqs.SendMessageOutput, error) {
	queueURLOutput, err := s.GetQueueURL(ctx, queueName)
	if err != nil {
		return nil, err
	}

	msgSQS, err := json.Marshal(body)
	if err != nil {
		logrus.Error("error on marshling a message")
		return nil, err
	}

	input := &sqs.SendMessageInput{
		DelaySeconds: delaySeconds,
		MessageBody:  aws.String(string(msgSQS)),
		QueueUrl:     queueURLOutput.QueueUrl,
	}

	return Client.SendMessage(ctx, input)
}

func (s service) GetQueueURL(ctx context.Context, queueName string) (*sqs.GetQueueUrlOutput, error) {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	return Client.GetQueueUrl(ctx, input)
}
