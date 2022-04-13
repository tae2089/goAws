package cloudwatch

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatch/types"
	"log"
	"time"
)

type CloudwatchInfo struct {
	AwsRegion       string
	AwsAccessKey    string
	AwsSecretKey    string
	AwsProfileName  string
	AwsInstanceId   string
	AwsInstanceName string
	AwsNamespace    string
	Client          *cloudwatch.Client
}

func (c *CloudwatchInfo) SetConfigByDefault() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(c.AwsRegion),
	)
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	c.Client = cloudwatch.NewFromConfig(cfg)
	return nil
}

// profile Name을 활용해서 Client 생성
func (c *CloudwatchInfo) SetConfigByProfile() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(c.AwsRegion),
		config.WithSharedConfigProfile(c.AwsProfileName))
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	c.Client = cloudwatch.NewFromConfig(cfg)
	return nil
}

//key를 활용해서 Client 생성
func (c *CloudwatchInfo) SetConfigByKey() error {
	creds := credentials.NewStaticCredentialsProvider(c.AwsAccessKey, c.AwsSecretKey, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(c.AwsRegion),
	)
	if err != nil {
		log.Printf("error: %v", err)
		//panic(err)
		return errors.New(err.Error())
	}
	c.Client = cloudwatch.NewFromConfig(cfg)
	return nil
}

func (c *CloudwatchInfo) PutMetricData(matrix ...CloudWatchMatrix) (out *cloudwatch.PutMetricDataOutput, err error) {
	fmt.Println("start PutMetricData")
	currentTime := time.Now()
	MetricData := []types.MetricDatum{}
	for _, watchMatrix := range matrix {
		metricName, metricValue := watchMatrix()
		metricDatum := types.MetricDatum{
			MetricName: aws.String(metricName),
			Value:      aws.Float64(metricValue),
			Timestamp:  &currentTime,
			Dimensions: []types.Dimension{
				{
					Name:  aws.String("instanceName"),
					Value: aws.String(c.AwsInstanceName),
				},
				{
					Name:  aws.String("instanceId"),
					Value: aws.String(c.AwsInstanceId),
				},
			},
		}
		MetricData = append(MetricData, metricDatum)
	}
	out, err = c.Client.PutMetricData(context.TODO(), &cloudwatch.PutMetricDataInput{
		MetricData: MetricData,
		Namespace:  aws.String(c.AwsNamespace),
	})
	fmt.Println("end PutMetricData")
	return
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
