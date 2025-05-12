package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	endpointURL = "https://pvos-tr-ist-01.portvmind.com"
	accessKeyID = "<ACCESS_KEY_ID>"
	secretKey   = "<SECRET_KEY>"
	region      = "tr-ist-01"
	bucketName  = "test-bucket-go"
	objectKey   = "test.txt"
	objectBody  = "This is a test file content for Ceph S3 API testing."
)

func main() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretKey, "")),
	)
	if err != nil {
		log.Fatalf("Unable to load SDK config: %v", err)
	}

	// Create S3 client with custom endpoint
	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(endpointURL)
	})

	// Run tests
	fmt.Println("Starting Ceph S3 API Test...")

	// Test 1: List buckets
	fmt.Println("\nTest 1: Listing buckets")
	if err := listBuckets(client); err != nil {
		log.Fatalf("Failed to list buckets: %v", err)
	}

	// Test 2: Create bucket
	fmt.Println("\nTest 2: Creating bucket")
	if err := createBucket(client, bucketName); err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	// Test 3: Upload object
	fmt.Println("\nTest 3: Uploading object")
	if err := uploadObject(client, bucketName, objectKey, objectBody); err != nil {
		log.Fatalf("Failed to upload object: %v", err)
	}

	// Test 4: Download and verify object
	fmt.Println("\nTest 4: Downloading and verifying object")
	if err := downloadAndVerifyObject(client, bucketName, objectKey, objectBody); err != nil {
		log.Fatalf("Failed to download and verify object: %v", err)
	}

	// Test 5: Delete object
	fmt.Println("\nTest 5: Deleting object")
	if err := deleteObject(client, bucketName, objectKey); err != nil {
		log.Fatalf("Failed to delete object: %v", err)
	}

	fmt.Println("\nAll tests completed successfully!")
}

func listBuckets(client *s3.Client) error {
	result, err := client.ListBuckets(context.TODO(), &s3.ListBucketsInput{})
	if err != nil {
		return err
	}

	fmt.Println("Available buckets:")
	for _, bucket := range result.Buckets {
		fmt.Printf("- %s (Created: %s)\n", *bucket.Name, bucket.CreationDate)
	}
	return nil
}

func createBucket(client *s3.Client, bucketName string) error {
	_, err := client.CreateBucket(context.TODO(), &s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		if strings.Contains(err.Error(), "BucketAlreadyOwnedByYou") {
			fmt.Printf("Bucket already exists: %s\n", bucketName)
			return nil
		}
		return err
	}

	fmt.Printf("Bucket successfully created: %s\n", bucketName)
	return nil
}

func uploadObject(client *s3.Client, bucketName, objectKey, objectBody string) error {
	_, err := client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
		Body:   strings.NewReader(objectBody),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Object successfully uploaded: %s\n", objectKey)
	return nil
}

func downloadAndVerifyObject(client *s3.Client, bucketName, objectKey, expectedContent string) error {
	result, err := client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		return err
	}

	content := string(body)
	fmt.Printf("Object successfully downloaded: %s\n", objectKey)
	fmt.Printf("Object size: %d bytes\n", len(content))
	fmt.Printf("Object content: %s\n", content)

	// Verify content
	if content == expectedContent {
		fmt.Println("Content verification: SUCCESS")
	} else {
		fmt.Println("Content verification: FAILED")
		fmt.Printf("Expected: %s\n", expectedContent)
		fmt.Printf("Actual: %s\n", content)
	}

	return nil
}

func deleteObject(client *s3.Client, bucketName, objectKey string) error {
	_, err := client.DeleteObject(context.TODO(), &s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	if err != nil {
		return err
	}

	fmt.Printf("Object successfully deleted: %s\n", objectKey)
	return nil
}
