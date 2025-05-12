# Portvmind Object Storage Examples

This repository contains example implementations for interacting with Portvmind Object Storage using different programming languages and their respective AWS SDKs.

## Overview

The examples demonstrate how to use AWS SDKs to interact with Portvmind Object Storage's S3 API. Each example includes:
- S3 API Authentication
- Bucket listing
- File upload
- File download
- Content verification
- File deletion

## Available Examples

### AWS CLI Example
Located in `examples/aws-cli/`
- Uses AWS CLI version 2.x (tested with aws-cli/2.26.6 Python/3.13.3 Darwin/23.2.0 source/arm64)
- Includes both basic and advanced S3 operations
- Provides a complete test script
- Requires AWS CLI 2.0.0 or higher

### Java Example
Located in `examples/java/ceph-s3-client/`
- Uses AWS SDK for Java 2.x
- Maven-based project
- Requires Java 11 or higher

### JavaScript Example
Located in `examples/javascript/`
- Uses AWS SDK for JavaScript v3
- Node.js-based project
- Requires Node.js 14 or higher

### Go Example
Located in `examples/golang/ceph-s3-client/`
- Uses AWS SDK for Go v2
- Go modules-based project
- Requires Go 1.16 or higher

## Prerequisites

- Portvmind Object Storage with S3 API enabled
- S3 API credentials (access key and secret key)
- Required development tools for each language:
  - AWS CLI 2.0.0 or higher
  - JDK 11 or higher
  - Node.js 14 or higher
  - Go 1.16 or higher

## Configuration

Each example requires the following configuration:
- Endpoint URL (e.g., https://pvos-tr-ist-01.portvmind.com)
- Access Key ID
- Secret Key
- Region (usually "default" for Portvmind Object Storage)
- Bucket name
- Object key
- Object content

Update these values in the respective example's configuration file before running.

## Running the Examples

### AWS CLI Example
```bash
cd examples/aws-cli
chmod +x test-s3-operations.sh
./test-s3-operations.sh
```

### Java Example
```bash
cd examples/java/ceph-s3-client
mvn clean install
mvn exec:java -Dexec.mainClass="com.example.ceph.Main"
```

### JavaScript Example
```bash
cd examples/javascript
npm install
npm start
```

### Go Example
```bash
cd examples/golang/ceph-s3-client
go mod tidy
go run main.go
```

## Error Handling

All examples include comprehensive error handling for S3 operations. If any operation fails, the program will log the error and exit with a non-zero status code.

## Contributing

Feel free to contribute to this repository by:
1. Adding new language examples
2. Improving existing examples
3. Adding more features or test scenarios
4. Fixing bugs or issues

## License

This project is licensed under the MIT License - see the LICENSE file for details.
