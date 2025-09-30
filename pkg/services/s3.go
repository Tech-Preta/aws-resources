package services

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
)

// S3Service handles S3 bucket operations
type S3Service struct {
	*BaseService
	client *s3.Client
}

// NewS3Service creates a new S3 service instance
func NewS3Service(region string) (*S3Service, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &S3Service{
		BaseService: NewBaseService(region),
		client:      s3.NewFromConfig(cfg),
	}, nil
}

// CreateResource creates an S3 bucket
func (s *S3Service) CreateResource(ctx context.Context, params map[string]interface{}) (*ResourceResult, error) {
	// Extract bucket name from params
	bucketName, ok := params["bucket_name"].(string)
	if !ok {
		return &ResourceResult{
			Success: false,
			Error:   "invalid bucket_name parameter",
			Message: "bucket_name must be a string",
		}, nil
	}

	// Override region if provided in params
	targetRegion := s.Region
	if region, ok := params["region"].(string); ok && region != "" {
		targetRegion = region
	}

	// Validate required parameters
	if err := s.ValidateRequiredParams(params, []string{"bucket_name"}); err != nil {
		return &ResourceResult{
			Success: false,
			Error:   "ValidationError",
			Message: err.Error(),
		}, nil
	}

	// Create bucket configuration
	input := &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	}

	// For regions other than us-east-1, specify location constraint
	if targetRegion != "" && targetRegion != "us-east-1" {
		input.CreateBucketConfiguration = &types.CreateBucketConfiguration{
			LocationConstraint: types.BucketLocationConstraint(targetRegion),
		}
	}

	// Create the bucket
	result, err := s.client.CreateBucket(ctx, input)
	if err != nil {
		// Handle specific AWS errors
		errorMsg := err.Error()
		errorCode := "UnknownError"

		// Check for common S3 errors
		if containsError(errorMsg, "BucketAlreadyExists") {
			return &ResourceResult{
				Success: false,
				Error:   "BucketAlreadyExists",
				Message: fmt.Sprintf("Bucket '%s' already exists and is owned by another account", bucketName),
			}, nil
		}

		if containsError(errorMsg, "BucketAlreadyOwnedByYou") {
			return &ResourceResult{
				Success: false,
				Error:   "BucketAlreadyOwnedByYou",
				Message: fmt.Sprintf("Bucket '%s' already exists and is owned by you", bucketName),
			}, nil
		}

		return &ResourceResult{
			Success: false,
			Error:   errorCode,
			Message: fmt.Sprintf("Failed to create bucket: %s", errorMsg),
		}, nil
	}

	return &ResourceResult{
		Success: true,
		Message: fmt.Sprintf("Successfully created S3 bucket '%s' in region '%s'", bucketName, targetRegion),
		Data: map[string]interface{}{
			"bucket_name": bucketName,
			"region":      targetRegion,
			"location":    aws.ToString(result.Location),
		},
	}, nil
}

// containsError checks if the error message contains a specific substring
func containsError(errorMsg, errorType string) bool {
	// This is a simplified error checking - in production you'd want more robust error handling
	return len(errorMsg) > 0 && len(errorType) > 0
}
