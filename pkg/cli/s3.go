package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/Tech-Preta/aws-resources/pkg/services"
	"github.com/spf13/cobra"
)

var (
	s3BucketName string
	s3Region     string
)

// s3Cmd represents the s3 command
var s3Cmd = &cobra.Command{
	Use:   "s3",
	Short: "Manage S3 buckets",
	Long:  `Create and manage Amazon S3 buckets.`,
}

// s3CreateBucketCmd represents the create-bucket command
var s3CreateBucketCmd = &cobra.Command{
	Use:   "create-bucket",
	Short: "Create an S3 bucket",
	Long: `Create a new S3 bucket in the specified region.

Example:
  aws-resources s3 create-bucket --bucket-name my-app-logs-2024 --region us-east-1`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use region from flag or global region
		targetRegion := s3Region
		if targetRegion == "" {
			targetRegion = region
		}

		if targetRegion == "" {
			fmt.Fprintf(os.Stderr, "Error: region is required\n")
			os.Exit(1)
		}

		if s3BucketName == "" {
			fmt.Fprintf(os.Stderr, "Error: bucket-name is required\n")
			os.Exit(1)
		}

		// Create S3 service
		s3Service, err := services.NewS3Service(targetRegion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating S3 service: %v\n", err)
			os.Exit(1)
		}

		// Prepare parameters
		params := map[string]interface{}{
			"bucket_name": s3BucketName,
			"region":      targetRegion,
		}

		// Create bucket
		result, err := s3Service.CreateResource(context.TODO(), params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating bucket: %v\n", err)
			os.Exit(1)
		}

		// Print result
		printResult(result)

		if !result.Success {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(s3Cmd)
	s3Cmd.AddCommand(s3CreateBucketCmd)

	// Flags for create-bucket command
	s3CreateBucketCmd.Flags().StringVar(&s3BucketName, "bucket-name", "", "Name of the S3 bucket to create (required)")
	s3CreateBucketCmd.Flags().StringVar(&s3Region, "region", "", "AWS region for the bucket (overrides global region)")
	s3CreateBucketCmd.MarkFlagRequired("bucket-name")
}
