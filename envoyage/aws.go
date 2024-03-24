package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func getAwsInstances() (map[int]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})
	instanceList := map[int]string{}

	if err != nil {
		fmt.Println("Error creating session:", err)
		return instanceList, err
	}
	svc := ec2.New(sess)

	result, err := svc.DescribeInstances(nil)
	if err != nil {
		fmt.Println("Error", err)
		return instanceList, err
	}
	for idx, res := range result.Reservations {
		for i, inst := range res.Instances {
			instanceList[i] = *inst.InstanceId
			fmt.Printf("%d. Instance ID: %s\n", idx+1, *inst.InstanceId)
		}
	}
	return instanceList, nil
}
