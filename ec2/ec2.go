package ec2

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"log"
)

type Ec2Info struct {
	Client *ec2.Client
}

func (e *Ec2Info) SetConfigByDefault() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	e.Client = ec2.NewFromConfig(cfg)
	return nil
}

func (e *Ec2Info) GetInstances() {
	result, err := e.Client.DescribeInstances(context.TODO(), &ec2.DescribeInstancesInput{})
	if err != nil {
		panic(err)
	}
	println(result)
}
