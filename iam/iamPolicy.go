package iam

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func (i *IamInfo) CreatePolicy(policyName string, policyData []byte) {
	i.iamClient.CreatePolicy(context.TODO(), &iam.CreatePolicyInput{
		PolicyName:     aws.String(policyName),
		PolicyDocument: aws.String(string(policyData)),
	})
}

func (i *IamInfo) GetPolicyList() {
	paginator, err := i.iamClient.ListPolicies(context.TODO(), &iam.ListPoliciesInput{
		Scope: types.PolicyScopeTypeAll,
	})

	if err != nil {
		panic(err)
	}
	for _, policy := range paginator.Policies {
		fmt.Println(policy.PolicyName)

	}
}

func (i *IamInfo) DescriptionPolicy(policyName string) {
	policy, err := i.iamClient.GetPolicy(context.TODO(), &iam.GetPolicyInput{
		PolicyArn: aws.String(policyName),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(policy.Policy.PolicyName)
}

func (i *IamInfo) DeletePolicy(policyArn string) {
	i.iamClient.DeletePolicy(context.TODO(), &iam.DeletePolicyInput{
		PolicyArn: aws.String(policyArn),
	})
}
