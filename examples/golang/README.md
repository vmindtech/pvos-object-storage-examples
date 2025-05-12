# Ceph S3 API Go Client

This project demonstrates how to use AWS S3 SDK with Ceph Object Storage's S3 API.

## Features

- S3 API Authentication
- Bucket listing
- File upload
- File download
- Content verification
- File deletion

## Prerequisites

- Go 1.16 or higher
- Ceph Object Storage with S3 API enabled
- S3 API credentials (access key and secret key)

## Installation

1. Clone the project
2. Install dependencies:
```bash
go mod tidy
```

## Configuration

Update the following constants in `main.go` according to your Ceph installation:
- `endpointURL`: Ceph S3 endpoint URL
- `accessKeyID`: Your S3 access key ID
- `secretKey`: Your S3 secret key
- `region`: Your S3 region
- `bucketName`: Test bucket name
- `objectKey`: Test object key
- `objectBody`: Test object content

## Usage

Build and run the project:
```bash
go build
./ceph-s3-client
```

Or run directly:
```bash
go run main.go
```

## Test Scenarios

The program runs the following test scenarios:

1. List all available buckets
2. Create a test bucket
3. Upload a test file
4. Download and verify the test file content
5. Delete the test file

## Error Handling

The program includes comprehensive error handling for all S3 operations. If any operation fails, the program will log the error and exit with a non-zero status code. 