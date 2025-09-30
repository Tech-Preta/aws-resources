package services

import (
	"context"
	"fmt"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

// EC2Service handles EC2 instance operations
type EC2Service struct {
	*BaseService
	client *ec2.Client
}

// NewEC2Service creates a new EC2 service instance
func NewEC2Service(region string) (*EC2Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &EC2Service{
		BaseService: NewBaseService(region),
		client:      ec2.NewFromConfig(cfg),
	}, nil
}

// CreateResource creates EC2 instances
func (e *EC2Service) CreateResource(ctx context.Context, params map[string]interface{}) (*ResourceResult, error) {
	// Extract required parameters
	imageID, ok := params["image_id"].(string)
	if !ok {
		return &ResourceResult{
			Success: false,
			Error:   "invalid image_id parameter",
			Message: "image_id must be a string",
		}, nil
	}

	instanceType, ok := params["instance_type"].(string)
	if !ok {
		return &ResourceResult{
			Success: false,
			Error:   "invalid instance_type parameter",
			Message: "instance_type must be a string",
		}, nil
	}

	keyName, ok := params["key_name"].(string)
	if !ok {
		return &ResourceResult{
			Success: false,
			Error:   "invalid key_name parameter",
			Message: "key_name must be a string",
		}, nil
	}

	// Handle count parameter - default to 1
	count := 1
	if countParam, ok := params["count"]; ok {
		switch v := countParam.(type) {
		case int:
			count = v
		case string:
			if parsed, err := strconv.Atoi(v); err == nil {
				count = parsed
			}
		}
	}

	// Validate count
	if count < 1 {
		return &ResourceResult{
			Success: false,
			Error:   "ValidationError",
			Message: "Count must be at least 1",
		}, nil
	}

	// Override region if provided in params
	targetRegion := e.Region
	if region, ok := params["region"].(string); ok && region != "" {
		targetRegion = region

		// Update client configuration for the new region
		cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(targetRegion))
		if err != nil {
			return &ResourceResult{
				Success: false,
				Error:   "ConfigurationError",
				Message: fmt.Sprintf("Failed to configure AWS client for region %s: %s", targetRegion, err.Error()),
			}, nil
		}
		e.client = ec2.NewFromConfig(cfg)
	}

	// Validate required parameters
	if err := e.ValidateRequiredParams(params, []string{"image_id", "instance_type", "key_name"}); err != nil {
		return &ResourceResult{
			Success: false,
			Error:   "ValidationError",
			Message: err.Error(),
		}, nil
	}

	// Create RunInstances input
	input := &ec2.RunInstancesInput{
		ImageId:      aws.String(imageID),
		MinCount:     aws.Int32(int32(count)),
		MaxCount:     aws.Int32(int32(count)),
		InstanceType: types.InstanceType(instanceType),
		KeyName:      aws.String(keyName),
	}

	// Launch instances
	result, err := e.client.RunInstances(ctx, input)
	if err != nil {
		errorMsg := err.Error()
		errorCode := "UnknownError"

		// Handle specific EC2 errors
		if containsError(errorMsg, "InvalidAMIID") {
			return &ResourceResult{
				Success: false,
				Error:   "InvalidAMIID",
				Message: fmt.Sprintf("Invalid AMI ID: %s", imageID),
			}, nil
		}

		if containsError(errorMsg, "InvalidKeyPair") {
			return &ResourceResult{
				Success: false,
				Error:   "InvalidKeyPair",
				Message: fmt.Sprintf("Invalid key pair: %s", keyName),
			}, nil
		}

		return &ResourceResult{
			Success: false,
			Error:   errorCode,
			Message: fmt.Sprintf("Failed to launch instances: %s", errorMsg),
		}, nil
	}

	// Extract instance information
	instances := make([]map[string]interface{}, len(result.Instances))
	for i, instance := range result.Instances {
		instances[i] = map[string]interface{}{
			"instance_id":   aws.ToString(instance.InstanceId),
			"state":         string(instance.State.Name),
			"image_id":      aws.ToString(instance.ImageId),
			"instance_type": string(instance.InstanceType),
			"key_name":      aws.ToString(instance.KeyName),
		}
	}

	return &ResourceResult{
		Success: true,
		Message: fmt.Sprintf("Successfully launched %d EC2 instance(s) in region '%s'", count, targetRegion),
		Data: map[string]interface{}{
			"instances":     instances,
			"region":        targetRegion,
			"image_id":      imageID,
			"instance_type": instanceType,
			"key_name":      keyName,
			"count":         count,
		},
	}, nil
}
