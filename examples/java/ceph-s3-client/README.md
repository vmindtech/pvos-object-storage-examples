# Ceph S3 API Java Client

This project demonstrates how to use AWS SDK with Ceph Object Storage's S3 API in Java.

## Features

- S3 API Authentication
- Bucket listing
- File upload
- File download
- Content verification
- File deletion

## Prerequisites

- Java 11 or higher
- Maven
- Ceph Object Storage with S3 API enabled
- S3 API credentials (access key and secret key)

## Installation

1. Clone the project
2. Install dependencies:
```bash
mvn clean install
```

## Configuration

Update the following constants in `Main.java` according to your Ceph installation:
- `ENDPOINT_URL`: Ceph S3 endpoint URL
- `ACCESS_KEY_ID`: Your S3 access key ID
- `SECRET_KEY`: Your S3 secret key
- `REGION`: Your S3 region
- `BUCKET_NAME`: Test bucket name
- `OBJECT_KEY`: Test object key
- `OBJECT_BODY`: Test object content

## Usage

Run the project:
```bash
mvn exec:java -Dexec.mainClass="com.example.ceph.Main"
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