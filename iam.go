package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"log"
)

type IamInfo struct {
	AwsS3Region    string
	AwsAccessKey   string
	AwsSecretKey   string
	AwsProfileName string
	iamClient      *iam.Client
}

func (i *IamInfo) SetConfigByDefault() error {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}

	i.iamClient = iam.NewFromConfig(cfg)

	return nil
}

// profile Name을 활용해서 Client 생성
func (i *IamInfo) SetConfigByProfile() error {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(i.AwsS3Region),
		config.WithSharedConfigProfile(i.AwsProfileName))
	if err != nil {
		log.Fatal(err)
		return errors.New(err.Error())
	}
	i.iamClient = iam.NewFromConfig(cfg)
	return nil
}

//key를 활용해서 Client 생성
func (i *IamInfo) SetConfigByKey() error {
	creds := credentials.NewStaticCredentialsProvider(i.AwsAccessKey, i.AwsSecretKey, "")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(creds),
		config.WithRegion(i.AwsS3Region),
	)
	if err != nil {
		log.Printf("error: %v", err)
		//panic(err)
		return errors.New(err.Error())
	}
	i.iamClient = iam.NewFromConfig(cfg)
	return nil
}

func (i *IamInfo) CreateGroup(groupName, path string) {
	output, err := i.iamClient.CreateGroup(context.TODO(), &iam.CreateGroupInput{
		GroupName: aws.String(groupName),
		Path:      aws.String(path),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(output.Group)
}

func (i *IamInfo) AddUserInGroup(groupName, userName string) {
	i.iamClient.GetGroup(context.TODO(), &iam.GetGroupInput{
		GroupName: aws.String(groupName),
	})
}

func (i *IamInfo) CreateLoginProfile(userName, password string) {
	i.iamClient.CreateLoginProfile(context.TODO(), &iam.CreateLoginProfileInput{
		UserName:              aws.String(userName),
		Password:              aws.String(password),
		PasswordResetRequired: true,
	})
}
