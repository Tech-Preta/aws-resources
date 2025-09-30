package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Tech-Preta/aws-resources/pkg/services"
	"github.com/spf13/cobra"
)

var (
	verbose bool
	region  string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "aws-resources",
	Short: "A CLI tool for managing AWS resources",
	Long: `AWS Resources CLI Tool - A command-line tool for managing AWS resources.

This tool provides a simple interface for creating and managing AWS resources
like S3 buckets and EC2 instances using the AWS SDK for Go v2.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	rootCmd.PersistentFlags().StringVarP(&region, "region", "r", "", "AWS region")
}

// printResult prints the result of an operation
func printResult(result *services.ResourceResult) {
	if verbose {
		// Print full JSON result in verbose mode
		data, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			fmt.Printf("Error marshaling result: %v\n", err)
			return
		}
		fmt.Println(string(data))
	} else {
		// Print simple success/error message
		if result.Success {
			fmt.Printf("✅ %s\n", result.Message)
		} else {
			fmt.Printf("❌ %s\n", result.Message)
			if result.Error != "" {
				fmt.Printf("   Error: %s\n", result.Error)
			}
		}
	}
}
