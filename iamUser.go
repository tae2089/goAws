package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

//User 생성해주기
func (i *IamInfo) CreateUser(userName, path string) {
	i.iamClient.CreateUser(context.TODO(), &iam.CreateUserInput{
		UserName: aws.String(userName),

		Path: aws.String(path),
		Tags: []types.Tag{
			types.Tag{Key: aws.String("Name"), Value: aws.String(userName)},
		},
	})
}

//User 정보 얻기
func (i *IamInfo) GetUser(userName string) {
	user, err := i.iamClient.GetUser(context.TODO(), &iam.GetUserInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user.User.UserId)
}

//User 삭제하기
func (i *IamInfo) DeleteUser(userName string) {
	user, err := i.iamClient.DeleteUser(context.TODO(), &iam.DeleteUserInput{
		UserName: aws.String(userName),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
