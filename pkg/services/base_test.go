package services

import (
	"testing"
)

func TestNewBaseService(t *testing.T) {
	region := "us-east-1"
	service := NewBaseService(region)

	if service.Region != region {
		t.Errorf("Expected region %s, got %s", region, service.Region)
	}
}

func TestValidateRequiredParams(t *testing.T) {
	service := NewBaseService("us-east-1")

	// Test successful validation
	params := map[string]interface{}{
		"bucket_name": "test-bucket",
		"region":      "us-east-1",
	}
	required := []string{"bucket_name"}

	err := service.ValidateRequiredParams(params, required)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	// Test missing parameter
	params = map[string]interface{}{}
	required = []string{"bucket_name"}

	err = service.ValidateRequiredParams(params, required)
	if err == nil {
		t.Error("Expected error for missing parameters")
	}

	// Test empty string parameter
	params = map[string]interface{}{
		"bucket_name": "",
	}
	required = []string{"bucket_name"}

	err = service.ValidateRequiredParams(params, required)
	if err == nil {
		t.Error("Expected error for empty string parameter")
	}

	// Test nil parameter
	params = map[string]interface{}{
		"bucket_name": nil,
	}
	required = []string{"bucket_name"}

	err = service.ValidateRequiredParams(params, required)
	if err == nil {
		t.Error("Expected error for nil parameter")
	}
}
