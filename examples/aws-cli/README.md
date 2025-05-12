# Portvmind Object Storage AWS CLI Examples

This directory contains examples of using AWS CLI to interact with Portvmind Object Storage's S3 API.

## Prerequisites

- AWS CLI version 2.0.0 or higher (tested with aws-cli/2.26.6 Python/3.13.3 Darwin/23.2.0 source/arm64)
  ```bash
  # Check AWS CLI version
  aws --version
  ```
- Portvmind Object Storage credentials (access key and secret key)

## Configuration

1. Configure AWS CLI with your Portvmind Object Storage credentials:
```bash
aws configure
```

Enter the following information when prompted:
- AWS Access Key ID: Your Portvmind Object Storage access key
- AWS Secret Access Key: Your Portvmind Object Storage secret key
- Default region name: default
- Default output format: json

2. Create a custom endpoint configuration:
```bash
aws configure set s3.endpoint_url https://pvos-tr-ist-01.portvmind.com
```

## Basic Operations

### List Buckets
```bash
aws s3 ls --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Create a Bucket (Method 1 - s3 command)
```bash
aws s3 mb s3://test-bucket-cli --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Create a Bucket (Method 2 - s3api command)
```bash
aws s3api create-bucket \
    --bucket test-bucket-cli \
    --create-bucket-configuration LocationConstraint=default \
    --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Upload a File
```bash
# Create a test file
echo "This is a test file content for Portvmind Object Storage testing." > test.txt

# Upload the file
aws s3 cp test.txt s3://test-bucket-cli/test.txt --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### List Objects in a Bucket
```bash
aws s3 ls s3://test-bucket-cli --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Download a File
```bash
aws s3 cp s3://test-bucket-cli/test.txt downloaded-test.txt --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Delete an Object
```bash
aws s3 rm s3://test-bucket-cli/test.txt --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Delete a Bucket
```bash
aws s3 rb s3://test-bucket-cli --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

## Advanced Operations

### Upload a Directory
```bash
aws s3 sync ./local-directory s3://test-bucket-cli/remote-directory --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Download a Directory
```bash
aws s3 sync s3://test-bucket-cli/remote-directory ./local-directory --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Set Object ACL
```bash
aws s3api put-object-acl --bucket test-bucket-cli --key test.txt --acl public-read --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

### Get Object ACL
```bash
aws s3api get-object-acl --bucket test-bucket-cli --key test.txt --endpoint-url https://pvos-tr-ist-01.portvmind.com
```

## Script Examples

### Complete Test Script
Create a file named `test-s3-operations.sh`:
```bash
#!/bin/bash

# Configuration
ENDPOINT_URL="https://pvos-tr-ist-01.portvmind.com"
BUCKET_NAME="test-bucket-cli"
TEST_FILE="test.txt"
TEST_CONTENT="This is a test file content for Portvmind Object Storage testing."

echo "Starting Portvmind Object Storage Test..."

# Create test file
echo "$TEST_CONTENT" > $TEST_FILE

# Create bucket
echo "Creating bucket..."
aws s3 mb s3://$BUCKET_NAME --endpoint-url $ENDPOINT_URL

# Upload file
echo "Uploading file..."
aws s3 cp $TEST_FILE s3://$BUCKET_NAME/$TEST_FILE --endpoint-url $ENDPOINT_URL

# List objects
echo "Listing objects..."
aws s3 ls s3://$BUCKET_NAME --endpoint-url $ENDPOINT_URL

# Download file
echo "Downloading file..."
aws s3 cp s3://$BUCKET_NAME/$TEST_FILE downloaded-$TEST_FILE --endpoint-url $ENDPOINT_URL

# Verify content
echo "Verifying content..."
if [ "$(cat downloaded-$TEST_FILE)" = "$TEST_CONTENT" ]; then
    echo "Content verification: SUCCESS"
else
    echo "Content verification: FAILED"
fi

# Cleanup
echo "Cleaning up..."
aws s3 rm s3://$BUCKET_NAME/$TEST_FILE --endpoint-url $ENDPOINT_URL
aws s3 rb s3://$BUCKET_NAME --endpoint-url $ENDPOINT_URL
rm $TEST_FILE downloaded-$TEST_FILE

echo "Test completed!"
```

Make the script executable and run it:
```bash
chmod +x test-s3-operations.sh
./test-s3-operations.sh
```

## Troubleshooting

### Common Issues

1. **Endpoint URL Error**
   - Ensure the endpoint URL is correct
   - Check if the service is accessible

2. **Authentication Error**
   - Verify your access key and secret key
   - Check if the credentials are properly configured

3. **Bucket Already Exists**
   - Use a different bucket name
   - Delete the existing bucket first

4. **Permission Denied**
   - Check your IAM permissions
   - Verify bucket policies

### Debug Mode

Enable debug mode for detailed information:
```bash
aws s3 ls --debug --endpoint-url https://pvos-tr-ist-01.portvmind.com
``` 