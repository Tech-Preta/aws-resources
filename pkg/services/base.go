package services

import (
	"context"
	"fmt"
)

// AWSService represents the interface that all AWS services must implement
type AWSService interface {
	CreateResource(ctx context.Context, params map[string]interface{}) (*ResourceResult, error)
}

// ResourceResult represents the result of a resource creation operation
type ResourceResult struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data,omitempty"`
	Error   string                 `json:"error,omitempty"`
}

// BaseService provides common functionality for all AWS services
type BaseService struct {
	Region string
}

// NewBaseService creates a new base service instance
func NewBaseService(region string) *BaseService {
	return &BaseService{
		Region: region,
	}
}

// ValidateRequiredParams validates that all required parameters are provided
func (b *BaseService) ValidateRequiredParams(params map[string]interface{}, required []string) error {
	var missing []string

	for _, param := range required {
		if value, exists := params[param]; !exists || value == nil || value == "" {
			missing = append(missing, param)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required parameters: %v", missing)
	}

	return nil
}
