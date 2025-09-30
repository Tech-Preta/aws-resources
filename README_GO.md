# AWS Resources CLI Tool

A command-line interface for creating and managing AWS resources built with Go and aws-sdk-go-v2.

## Features

- **S3 Bucket Management**: Create S3 buckets with region specification
- **EC2 Instance Management**: Launch EC2 instances with customizable configuration
- **Modern Go Implementation**: Built with aws-sdk-go-v2 for optimal performance
- **Comprehensive Error Handling**: User-friendly AWS error messages
- **Professional CLI**: Built with Cobra for excellent user experience

## Prerequisites

- Go 1.21 or higher
- AWS credentials configured (via AWS CLI, environment variables, or IAM roles)

## Installation

### From Source

1. Clone this repository:
```bash
git clone https://github.com/Tech-Preta/aws-resources.git
cd aws-resources
```

2. Build the application:
```bash
make build
```

3. Install the binary (optional):
```bash
make install
```

### Configure AWS Credentials

Choose one method:
- Using AWS CLI: `aws configure`
- Using environment variables: `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
- Using IAM roles (if running on EC2)

## Usage

### S3 Operations

#### Create an S3 bucket:
```bash
./bin/aws-resources s3 create-bucket --bucket-name my-app-logs-2024 --region us-east-1
```

### EC2 Operations

#### Launch a single t2.micro instance:
```bash
./bin/aws-resources ec2 create-instances \
  --image-id ami-0abcdef1234567890 \
  --instance-type t2.micro \
  --key-name my-keypair \
  --region us-west-2
```

#### Launch multiple instances with verbose output:
```bash
./bin/aws-resources -v ec2 create-instances \
  --image-id ami-0abcdef1234567890 \
  --instance-type t3.small \
  --key-name production-key \
  --count 3 \
  --region eu-west-1
```

### Global Options

- `-v, --verbose`: Enable verbose output with full JSON responses
- `-r, --region`: Specify AWS region globally

## Development

### Building

```bash
make build          # Build the application
make dev            # Build with race detection
make clean          # Clean build artifacts
```

### Testing

```bash
make test           # Run tests
make test-coverage  # Run tests with coverage
make check          # Run format, vet, and test
```

### Code Quality

```bash
make fmt            # Format code
make vet            # Vet code
```

## Architecture

The CLI is built with extensibility in mind:

- **Base Service Interface**: Common functionality for all AWS services
- **Service-Specific Implementations**: Dedicated modules for each AWS service (S3, EC2)
- **Cobra CLI Framework**: Professional command-line interface with subcommands
- **aws-sdk-go-v2**: Modern, efficient AWS SDK for Go
- **Comprehensive Error Handling**: User-friendly error messages with proper AWS error handling

## Project Structure

```
├── cmd/aws-resources/     # Main application entry point
├── pkg/
│   ├── cli/              # CLI commands and interface
│   └── services/         # AWS service implementations
├── Makefile              # Build and development commands
├── go.mod               # Go module definition
└── go.sum              # Go module checksums
```

## Error Handling

The CLI provides comprehensive error handling:

- **Validation Errors**: Missing or invalid parameters
- **AWS Service Errors**: Specific handling for common AWS errors (e.g., BucketAlreadyExists)
- **Network Errors**: Connection and timeout issues
- **Permission Errors**: IAM and credential-related issues

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the GNU General Public License v3.0. See the [LICENSE](LICENSE) file for details.

---

Feito com ❤️ por [Natália Granato](https://github.com/nataliagranato).