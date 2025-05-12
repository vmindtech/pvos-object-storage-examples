package com.example.ceph;

import software.amazon.awssdk.auth.credentials.AwsBasicCredentials;
import software.amazon.awssdk.auth.credentials.StaticCredentialsProvider;
import software.amazon.awssdk.regions.Region;
import software.amazon.awssdk.services.s3.S3Client;
import software.amazon.awssdk.services.s3.model.*;

import java.net.URI;
import java.nio.charset.StandardCharsets;

public class Main {
    private static final String ENDPOINT_URL = "https://pvos-tr-ist-01.portvmind.com";
    private static final String ACCESS_KEY_ID = "<ACCESS_KEY_ID>";
    private static final String SECRET_KEY = "<SECRET_KEY>";
    private static final String REGION = "default";
    private static final String BUCKET_NAME = "test-bucket-java";
    private static final String OBJECT_KEY = "test.txt";
    private static final String OBJECT_BODY = "This is a test file content for Ceph S3 API testing.";

    public static void main(String[] args) {
        try {
            System.out.println("Starting Ceph S3 API Test...");

            // Create S3 client
            S3Client s3Client = S3Client.builder()
                    .endpointOverride(URI.create(ENDPOINT_URL))
                    .region(Region.of(REGION))
                    .credentialsProvider(StaticCredentialsProvider.create(
                            AwsBasicCredentials.create(ACCESS_KEY_ID, SECRET_KEY)))
                    .forcePathStyle(true)
                    .build();

            // Test 1: List buckets
            System.out.println("\nTest 1: Listing buckets");
            listBuckets(s3Client);

            // Test 2: Create bucket
            System.out.println("\nTest 2: Creating bucket");
            createBucket(s3Client);

            // Test 3: Upload object
            System.out.println("\nTest 3: Uploading object");
            uploadObject(s3Client);

            // Test 4: Download and verify object
            System.out.println("\nTest 4: Downloading and verifying object");
            downloadAndVerifyObject(s3Client);

            // Test 5: Delete object
            System.out.println("\nTest 5: Deleting object");
            deleteObject(s3Client);

            System.out.println("\nAll tests completed successfully!");
        } catch (Exception e) {
            System.err.println("An error occurred: " + e.getMessage());
            e.printStackTrace();
        }
    }

    private static void listBuckets(S3Client s3Client) {
        ListBucketsResponse response = s3Client.listBuckets();
        System.out.println("Available buckets:");
        response.buckets().forEach(bucket -> 
            System.out.printf("- %s (Created: %s)%n", bucket.name(), bucket.creationDate())
        );
    }

    private static void createBucket(S3Client s3Client) {
        try {
            CreateBucketRequest request = CreateBucketRequest.builder()
                    .bucket(BUCKET_NAME)
                    .createBucketConfiguration(CreateBucketConfiguration.builder()
                            .locationConstraint(REGION)
                            .build())
                    .build();
            s3Client.createBucket(request);
            System.out.printf("Bucket successfully created: %s%n", BUCKET_NAME);
        } catch (BucketAlreadyOwnedByYouException e) {
            System.out.printf("Bucket already exists: %s%n", BUCKET_NAME);
        }
    }

    private static void uploadObject(S3Client s3Client) {
        PutObjectRequest request = PutObjectRequest.builder()
                .bucket(BUCKET_NAME)
                .key(OBJECT_KEY)
                .build();

        s3Client.putObject(request, software.amazon.awssdk.core.sync.RequestBody.fromString(OBJECT_BODY));
        System.out.printf("Object successfully uploaded: %s%n", OBJECT_KEY);
    }

    private static void downloadAndVerifyObject(S3Client s3Client) {
        GetObjectRequest request = GetObjectRequest.builder()
                .bucket(BUCKET_NAME)
                .key(OBJECT_KEY)
                .build();

        String content = s3Client.getObjectAsBytes(request).asString(StandardCharsets.UTF_8);
        System.out.printf("Object successfully downloaded: %s%n", OBJECT_KEY);
        System.out.printf("Object size: %d bytes%n", content.length());
        System.out.printf("Object content: %s%n", content);

        // Verify content
        if (content.equals(OBJECT_BODY)) {
            System.out.println("Content verification: SUCCESS");
        } else {
            System.out.println("Content verification: FAILED");
            System.out.printf("Expected: %s%n", OBJECT_BODY);
            System.out.printf("Actual: %s%n", content);
        }
    }

    private static void deleteObject(S3Client s3Client) {
        DeleteObjectRequest request = DeleteObjectRequest.builder()
                .bucket(BUCKET_NAME)
                .key(OBJECT_KEY)
                .build();

        s3Client.deleteObject(request);
        System.out.printf("Object successfully deleted: %s%n", OBJECT_KEY);
    }
} 