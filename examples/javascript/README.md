# Ceph S3 API JavaScript Client

This project demonstrates how to use AWS SDK with Ceph Object Storage's S3 API in JavaScript.

## Features

- S3 API Authentication
- Bucket listing
- File upload
- File download
- Content verification
- File deletion

## Prerequisites

- Node.js 14 or higher
- npm or yarn
- Ceph Object Storage with S3 API enabled
- S3 API credentials (access key and secret key)

## Installation

1. Clone the project
2. Install dependencies:
```bash
npm install
```

## Configuration

Update the following constants in `index.js` according to your Ceph installation:
- `endpointURL`: Ceph S3 endpoint URL
- `accessKeyId`: Your S3 access key ID
- `secretAccessKey`: Your S3 secret key
- `region`: Your S3 region
- `bucketName`: Test bucket name
- `objectKey`: Test object key
- `objectBody`: Test object content

## Usage

Run the project:
```bash
npm start
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