package cli

import (
	"context"
	"fmt"
	"os"

	"github.com/Tech-Preta/aws-resources/pkg/services"
	"github.com/spf13/cobra"
)

var (
	ec2ImageID      string
	ec2InstanceType string
	ec2KeyName      string
	ec2Count        int
	ec2Region       string
)

// ec2Cmd represents the ec2 command
var ec2Cmd = &cobra.Command{
	Use:   "ec2",
	Short: "Manage EC2 instances",
	Long:  `Create and manage Amazon EC2 instances.`,
}

// ec2CreateInstancesCmd represents the create-instances command
var ec2CreateInstancesCmd = &cobra.Command{
	Use:   "create-instances",
	Short: "Create EC2 instances",
	Long: `Launch new EC2 instances with the specified configuration.

Example:
  aws-resources ec2 create-instances \
    --image-id ami-0abcdef1234567890 \
    --instance-type t2.micro \
    --key-name my-keypair \
    --region us-west-2

  aws-resources ec2 create-instances \
    --image-id ami-0abcdef1234567890 \
    --instance-type t3.small \
    --key-name production-key \
    --count 3 \
    --region eu-west-1`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use region from flag or global region
		targetRegion := ec2Region
		if targetRegion == "" {
			targetRegion = region
		}

		if targetRegion == "" {
			fmt.Fprintf(os.Stderr, "Error: region is required\n")
			os.Exit(1)
		}

		// Validate required parameters
		if ec2ImageID == "" {
			fmt.Fprintf(os.Stderr, "Error: image-id is required\n")
			os.Exit(1)
		}
		if ec2InstanceType == "" {
			fmt.Fprintf(os.Stderr, "Error: instance-type is required\n")
			os.Exit(1)
		}
		if ec2KeyName == "" {
			fmt.Fprintf(os.Stderr, "Error: key-name is required\n")
			os.Exit(1)
		}

		// Create EC2 service
		ec2Service, err := services.NewEC2Service(targetRegion)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating EC2 service: %v\n", err)
			os.Exit(1)
		}

		// Prepare parameters
		params := map[string]interface{}{
			"image_id":      ec2ImageID,
			"instance_type": ec2InstanceType,
			"key_name":      ec2KeyName,
			"count":         ec2Count,
			"region":        targetRegion,
		}

		// Create instances
		result, err := ec2Service.CreateResource(context.TODO(), params)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating instances: %v\n", err)
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
	rootCmd.AddCommand(ec2Cmd)
	ec2Cmd.AddCommand(ec2CreateInstancesCmd)

	// Flags for create-instances command
	ec2CreateInstancesCmd.Flags().StringVar(&ec2ImageID, "image-id", "", "AMI ID to launch the instance from (required)")
	ec2CreateInstancesCmd.Flags().StringVar(&ec2InstanceType, "instance-type", "", "EC2 instance type (e.g., t2.micro, t3.small) (required)")
	ec2CreateInstancesCmd.Flags().StringVar(&ec2KeyName, "key-name", "", "Name of the EC2 Key Pair for SSH access (required)")
	ec2CreateInstancesCmd.Flags().IntVar(&ec2Count, "count", 1, "Number of instances to launch")
	ec2CreateInstancesCmd.Flags().StringVar(&ec2Region, "region", "", "AWS region where instances will be launched (overrides global region)")

	ec2CreateInstancesCmd.MarkFlagRequired("image-id")
	ec2CreateInstancesCmd.MarkFlagRequired("instance-type")
	ec2CreateInstancesCmd.MarkFlagRequired("key-name")
}
