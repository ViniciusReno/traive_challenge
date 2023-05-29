// nolint
package aws

import (
	"context"
	"fmt"

	appConfig "github.com/ViniciusReno/traive/internal/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

func awsConfig(region, sqsEndpoint string) aws.Config {
	if sqsEndpoint == "" {

		cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
		if err != nil {
			fmt.Println(err)
		}

		return cfg
	}

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {

		if sqsEndpoint != "" && service == sqs.ServiceID {
			return aws.Endpoint{
				URL:           sqsEndpoint,
				SigningRegion: region,
			}, nil
		}

		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(customResolver))
	if err != nil {
		fmt.Println(err)
	}
	return cfg
}

// Session return the current aws session
func AWSSession() aws.Config {
	region := appConfig.Config("AWS_REGION")
	sqsEndpoint := appConfig.Config("AWS_SQS_ENDPOINT")

	return awsConfig(region, sqsEndpoint)
}
