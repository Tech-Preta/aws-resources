# AWS Resources CLI Tool

🚀 **Now powered by Go and aws-sdk-go-v2 with Bubble Tea TUI!**

A modern Terminal User Interface (TUI) for creating and managing AWS resources. This project has been refactored from Python to Go with an interactive interface powered by Bubble Tea.

![Main Menu](https://github.com/user-attachments/assets/8f724ca2-b911-4446-8f0a-986a881137d7)

## 🔄 Migration Notice

This project has been **migrated from Python (boto3) to Go (aws-sdk-go-v2)** with a modern **Bubble Tea TUI interface** to provide:
- ⚡ Better performance and lower memory usage
- 📦 Single binary deployment (no dependencies)
- 🔧 Modern AWS SDK with the latest features
- 🎨 Beautiful interactive Terminal User Interface
- 🖱️ Intuitive navigation with keyboard controls

**Legacy Python code is preserved** for reference, but the Go implementation with Bubble Tea TUI is now the primary version.

## Features

- 🎨 **Interactive TUI**: Beautiful terminal interface with keyboard navigation
- 🪣 **S3 Bucket Management**: Create S3 buckets with region specification
- 💻 **EC2 Instance Management**: Launch EC2 instances with customizable configuration
- 🚀 **Modern Go Implementation**: Built with aws-sdk-go-v2 for optimal performance
- 🛠️ **Comprehensive Error Handling**: User-friendly AWS error messages
- ⌨️ **Keyboard Navigation**: Arrow keys, Enter, Tab, and Esc for full control

## Prerequisites

- Go 1.21 or higher
- AWS credentials configured (via AWS CLI, environment variables, or IAM roles)

## Installation

### Quick Start

1. Clone and build:
```bash
git clone https://github.com/Tech-Preta/aws-resources.git
cd aws-resources
make build
```

2. Configure AWS credentials:
```bash
aws configure
# OR set environment variables:
# export AWS_ACCESS_KEY_ID=your_access_key
# export AWS_SECRET_ACCESS_KEY=your_secret_key
```

3. Run the interactive TUI:
```bash
./bin/aws-resources
```

## Usage

### Interactive TUI Navigation

The application provides a beautiful terminal user interface with the following controls:

- **↑/↓ or k/j**: Navigate menu options
- **Enter**: Select option or confirm action
- **Tab**: Switch between form fields
- **Esc**: Go back to previous screen
- **q**: Quit application

### S3 Operations

1. Select "S3 - Manage Buckets" from the main menu
2. Choose "Create Bucket"
3. Fill in:
   - Bucket Name: Enter your desired bucket name
   - Region: Specify AWS region (default: us-east-1)
4. Select "Create Bucket" to proceed

### EC2 Operations

1. Select "EC2 - Manage Instances" from the main menu
2. Choose "Create Instances"
3. Fill in:
   - Image ID (AMI): AMI ID to launch (e.g., ami-0abcdef1234567890)
   - Instance Type: EC2 instance type (e.g., t2.micro, t3.small)
   - Key Name: Name of your EC2 Key Pair for SSH access
   - Count: Number of instances to launch (default: 1)
   - Region: AWS region (default: us-east-1)
4. Select "Launch Instances" to proceed

## Development

### Build Commands

```bash
make build          # Build the application
make test           # Run tests
make clean          # Clean build artifacts
make check          # Run format, vet, and test
make help           # Show all available commands
```

## Migration from Command-Line Version

The new Bubble Tea TUI provides a more intuitive alternative to command-line arguments:

### Before (Command-line):
```bash
./bin/aws-resources s3 create-bucket --bucket-name mybucket --region us-east-1
./bin/aws-resources ec2 create-instances --image-id ami-123 --instance-type t2.micro --key-name mykey --region us-west-2
```

### After (Interactive TUI):
```bash
./bin/aws-resources
# Then navigate through the beautiful menu system
```

## Architecture

```
├── cmd/aws-resources/     # Main application entry point
├── pkg/
│   ├── cli/              # Bubble Tea TUI interface
│   └── services/         # AWS service implementations
├── Makefile              # Build and development commands
├── go.mod               # Go module definition
└── go.sum              # Go module checksums
```

## Technology Stack

- **Go 1.21+**: Modern, fast, and efficient programming language
- **aws-sdk-go-v2**: Latest AWS SDK for Go with best practices
- **Bubble Tea**: Powerful TUI framework for beautiful terminal interfaces
- **Lipgloss**: Styling library for elegant terminal output

## Legacy Python Implementation

The original Python implementation is preserved in this repository:
- `aws_cli/` - Original Python CLI implementation
- `aws_ui/` - Flask web interface (deprecated)
- `tests/` - Python unit tests

These files are kept for reference but are no longer actively maintained.

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

## Usage

### Modern Package-based CLI (Recommended for S3 & EC2)

#### S3 Bucket Creation

Create a new S3 bucket:

```bash
aws-resources s3 create-bucket --bucket-name my-unique-bucket --region us-west-2
```

**Parameters:**
- `--bucket-name`: Name of the S3 bucket (must be globally unique)
- `--region`: AWS region for the bucket (e.g., us-east-1, us-west-2)

#### EC2 Instance Creation

Launch new EC2 instances:

```bash
aws-resources ec2 create-instances \
  --image-id ami-12345678 \
  --instance-type t2.micro \
  --key-name my-ec2-key \
  --count 2 \
  --region us-east-1
```

**Parameters:**
- `--image-id`: AMI ID to launch the instance from (required)
- `--instance-type`: EC2 instance type like t2.micro, t3.small (required)
- `--key-name`: Name of the EC2 Key Pair for SSH access (required)
- `--count`: Number of instances to launch (default: 1)
- `--region`: AWS region for the instances (required)

#### Global Options

- `-v, --verbose`: Enable verbose output for detailed information

#### Help

Get help for any command:

```bash
aws-resources --help
aws-resources s3 --help
aws-resources ec2 create-instances --help
```

### Extended Service CLI (All AWS Services)

#### General Syntax
```bash
python aws_cli.py <service> <action> [options]
```

#### S3 Operations

Create an S3 bucket:
```bash
python aws_cli.py s3 create-bucket --name my-unique-bucket-name --region us-east-1
```

#### EC2 Operations

Create an EC2 instance:
```bash
python aws_cli.py ec2 create-instance --image-id ami-0abcdef1234567890 --instance-type t2.micro --region us-east-1
```

With optional parameters:
```bash
python aws_cli.py ec2 create-instance --image-id ami-0abcdef1234567890 --instance-type t2.micro --region us-east-1 --key-name my-key-pair --security-groups default
```

#### DynamoDB Operations

Create a simple table with partition key only:
```bash
python aws_cli.py dynamodb create-table --name my-table --partition-key id --region us-east-1
```

Create a table with both partition and sort keys:
```bash
python aws_cli.py dynamodb create-table --name my-table --partition-key userId --partition-key-type S --sort-key timestamp --sort-key-type N --region us-east-1
```

#### RDS Operations

Create an RDS MySQL instance:
```bash
python aws_cli.py rds create-instance --identifier my-database --engine mysql --username admin --password mypassword123 --region us-east-1
```

With additional options:
```bash
python aws_cli.py rds create-instance --identifier my-database --engine postgres --instance-class db.t3.small --username admin --password mypassword123 --storage 50 --publicly-accessible --region us-east-1
```

#### Lambda Operations

Create a Lambda function (with default Hello World code):
```bash
python aws_cli.py lambda create-function --name my-function --runtime python3.9 --role arn:aws:iam::123456789012:role/lambda-execution-role --region us-east-1
```

With custom code and settings:
```bash
python aws_cli.py lambda create-function --name my-function --runtime python3.9 --role arn:aws:iam::123456789012:role/lambda-execution-role --handler index.handler --code-file my-function.zip --timeout 60 --memory 256 --region us-east-1
```

#### SNS Operations

Create an SNS topic:
```bash
python aws_cli.py sns create-topic --name my-topic --region us-east-1
```

With display name:
```bash
python aws_cli.py sns create-topic --name my-topic --display-name "My Notification Topic" --region us-east-1
```

## Examples

### Modern CLI Examples

#### Create an S3 bucket in us-east-1:
```bash
aws-resources s3 create-bucket --bucket-name my-app-logs-2024 --region us-east-1
```

#### Launch a single t2.micro instance:
```bash
aws-resources ec2 create-instances \
  --image-id ami-0abcdef1234567890 \
  --instance-type t2.micro \
  --key-name my-keypair \
  --region us-west-2
```

#### Launch multiple instances with verbose output:
```bash
aws-resources -v ec2 create-instances \
  --image-id ami-0abcdef1234567890 \
  --instance-type t3.small \
  --key-name production-key \
  --count 3 \
  --region eu-west-1
```

## Architecture

The CLI is built with extensibility in mind:

- **Base Service Class**: Common functionality for all AWS services
- **Service-Specific Modules**: Dedicated modules for each AWS service
- **Argument Parsing**: Hierarchical subcommands using argparse
- **Error Handling**: Comprehensive error handling with user-friendly messages
## Command Reference (Extended CLI)

### Global Options
- `--region`: AWS region (required for all commands)

### S3 create-bucket
- `--name`: Bucket name (must be globally unique)

### EC2 create-instance
- `--image-id`: AMI ID to launch (required)
- `--instance-type`: Instance type (required, e.g., t2.micro)
- `--key-name`: EC2 Key Pair name (optional)
- `--security-groups`: Security group names (optional, space-separated)

### DynamoDB create-table
- `--name`: Table name (required)
- `--partition-key`: Partition key attribute name (required)
- `--partition-key-type`: Partition key type - S/N/B (default: S)
- `--sort-key`: Sort key attribute name (optional)
- `--sort-key-type`: Sort key type - S/N/B (default: S)

### RDS create-instance
- `--identifier`: DB instance identifier (required)
- `--engine`: Database engine - mysql/postgres/mariadb/oracle-ee/sqlserver-ex (required)
- `--username`: Master username (required)
- `--password`: Master password (required)
- `--instance-class`: DB instance class (default: db.t3.micro)
- `--storage`: Allocated storage in GB (default: 20)
- `--security-groups`: VPC security group IDs (optional, space-separated)
- `--publicly-accessible`: Make instance publicly accessible (flag)

### Lambda create-function
- `--name`: Function name (required)
- `--runtime`: Runtime environment (required) - python3.9/python3.10/python3.11/nodejs18.x/nodejs20.x/java11/dotnet6
- `--role`: IAM role ARN (required)
- `--handler`: Function handler (default: lambda_function.lambda_handler)
- `--code-file`: Path to code zip file (optional, creates Hello World if not provided)
- `--description`: Function description (optional)
- `--timeout`: Function timeout in seconds (optional)
- `--memory`: Memory size in MB (optional)

### SNS create-topic
- `--name`: Topic name (required)
- `--display-name`: Topic display name (optional)

## Requirements

- Python 3.8+
- boto3 >= 1.26.0
- Valid AWS credentials

## Error Handling

The tool provides clear error messages for common scenarios:
- Missing AWS credentials
- Resource already exists
- Invalid parameters
- AWS service errors

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the GPL-3.0 License - see the [LICENSE](LICENSE) file for details.

---

## Web Interface (New)

This project now includes a user-friendly web interface for creating AWS resources.

### Running the Web App

1.  **Install Flask**:
    ```bash
    pip install Flask
    ```

2.  **Run the application**:
    ```bash
    python -m aws_ui.app
    ```
    Or, if installed via `pip install -e .`:
    ```bash
    aws-resources-web
    ```

3.  Open your browser and navigate to `http://127.0.0.1:5000`.

### Web App Features
- Create S3 buckets
- Launch EC2 instances
- Real-time feedback on resource creation status

---

Feito com ❤️ por [Natália Granato](https://github.com/nataliagranato).
