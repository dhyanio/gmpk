package main
// Just run this `go run main.go <instance_ID>`
import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)
const (
	region = "us-east-1"
)
type MetaData struct {
	BlockDeviceMappings BlockDevice `json:"blockDeviceMappings"`
	ImageId             string      `json:"imageId"`
	InstanceId          string      `json:"instanceId"`
	InstanceType        string      `json:"instanceType"`
	KeyName             string      `json:"keyName"`
	Platform            string      `json:"platform"`
	PrivateIpAddress    string      `json:"privateIpAddress"`
	PublicIpAddress     string      `json:"publicIpAddress"`
	SubnetId            string      `json:"subnetId"`
	VpcId               string      `json:"vpcId"`
}
type BlockDevice struct {
	VolumeId   string `json:"volumeId"`
	DeviceName string `json:"deviceName"`
	Status     string `json:"status"`
}
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
func fetchInstanceMetadata(instanceID string) {
	var jsonData MetaData
	awsSess, err1 := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err1 != nil {
		exitErrorf("Not able to create aws session, %v", err1)
	}
	svc := ec2.New(awsSess)
	instanceInput := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{aws.String(instanceID)},
	}
	result, err2 := svc.DescribeInstances(instanceInput)
	if err2 != nil {
		exitErrorf("Not able to describe EC2 instance, %v", err2)
	}
	for _, value1 := range result.Reservations[0].Instances[0].BlockDeviceMappings {
		jsonData.BlockDeviceMappings.DeviceName = *value1.DeviceName
		jsonData.BlockDeviceMappings.Status = *value1.Ebs.Status
		jsonData.BlockDeviceMappings.VolumeId = *value1.Ebs.VolumeId
	}
	for _, value2 := range result.Reservations[0].Instances {
		jsonData.ImageId = *value2.ImageId
		jsonData.InstanceId = *value2.InstanceId
		jsonData.InstanceType = *value2.InstanceType
		jsonData.KeyName = *value2.KeyName
		jsonData.Platform = *value2.Platform
		jsonData.PrivateIpAddress = *value2.PrivateIpAddress
		jsonData.PublicIpAddress = *value2.PublicIpAddress
		jsonData.SubnetId = *value2.SubnetId
		jsonData.VpcId = *value2.VpcId
	}
	jsonOutput, err3 := json.MarshalIndent(jsonData, "", "\t")
	if err1 != nil {
		exitErrorf("Not able to create json data, %v", err3)
	}
	fmt.Print(string(jsonOutput))
}
func main() {
	EC2InstanceID := os.Args[1]
	fetchInstanceMetadata(EC2InstanceID)
}
