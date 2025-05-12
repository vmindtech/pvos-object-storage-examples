#!/bin/bash

# Configuration
ENDPOINT_URL="https://pvos-tr-ist-01.portvmind.com"
BUCKET_NAME="test-bucket-cli"
TEST_FILE="test.txt"
TEST_CONTENT="This is a test file content for Portvmind Object Storage testing."

echo "Starting Portvmind Object Storage Test..."

# Create test file
echo "$TEST_CONTENT" > $TEST_FILE

# Create bucket (Method 1 - s3 command)
echo "Creating bucket using s3 command..."
aws s3 mb s3://$BUCKET_NAME --endpoint-url $ENDPOINT_URL

# Alternative: Create bucket (Method 2 - s3api command)
# echo "Creating bucket using s3api command..."
# aws s3api create-bucket \
#     --bucket $BUCKET_NAME \
#     --create-bucket-configuration LocationConstraint=default \
#     --endpoint-url $ENDPOINT_URL

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