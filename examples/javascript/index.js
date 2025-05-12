import { S3Client, ListBucketsCommand, CreateBucketCommand, PutObjectCommand, GetObjectCommand, DeleteObjectCommand } from "@aws-sdk/client-s3";
import { fromEnv } from "@aws-sdk/credential-providers";

const endpointURL = "https://pvos-tr-ist-01.portvmind.com";
const accessKeyId = "<ACCESS_KEY_ID>";
const secretAccessKey = "<SECRET_KEY>";
const region = "default";
const bucketName = "test-bucket-js";
const objectKey = "test.txt";
const objectBody = "This is a test file content for Ceph S3 API testing.";

// Create S3 client
const s3Client = new S3Client({
    endpoint: endpointURL,
    region: region,
    credentials: {
        accessKeyId: accessKeyId,
        secretAccessKey: secretAccessKey
    },
    forcePathStyle: true,
    disableHostPrefix: true
});

async function main() {
    try {
        console.log("Starting Ceph S3 API Test...");

        // Test 1: List buckets
        console.log("\nTest 1: Listing buckets");
        await listBuckets();

        // Test 2: Create bucket
        console.log("\nTest 2: Creating bucket");
        await createBucket();

        // Test 3: Upload object
        console.log("\nTest 3: Uploading object");
        await uploadObject();

        // Test 4: Download and verify object
        console.log("\nTest 4: Downloading and verifying object");
        await downloadAndVerifyObject();

        // Test 5: Delete object
        console.log("\nTest 5: Deleting object");
        await deleteObject();

        console.log("\nAll tests completed successfully!");
    } catch (error) {
        console.error("An error occurred:", error);
    }
}

async function listBuckets() {
    const command = new ListBucketsCommand({});
    const response = await s3Client.send(command);
    
    console.log("Available buckets:");
    response.Buckets.forEach(bucket => {
        console.log(`- ${bucket.Name} (Created: ${bucket.CreationDate})`);
    });
}

async function createBucket() {
    try {
        const command = new CreateBucketCommand({
            Bucket: bucketName,
            CreateBucketConfiguration: {
                LocationConstraint: "default"
            }
        });
        await s3Client.send(command);
        console.log(`Bucket successfully created: ${bucketName}`);
    } catch (error) {
        if (error.name === "BucketAlreadyOwnedByYou") {
            console.log(`Bucket already exists: ${bucketName}`);
        } else {
            throw error;
        }
    }
}

async function uploadObject() {
    const command = new PutObjectCommand({
        Bucket: bucketName,
        Key: objectKey,
        Body: objectBody
    });
    
    const response = await s3Client.send(command);
    console.log(`Object successfully uploaded: ${objectKey}`);
    console.log(`ETag: ${response.ETag}`);
}

async function downloadAndVerifyObject() {
    const command = new GetObjectCommand({
        Bucket: bucketName,
        Key: objectKey
    });
    
    const response = await s3Client.send(command);
    const content = await response.Body.transformToString();
    
    console.log(`Object successfully downloaded: ${objectKey}`);
    console.log(`Object size: ${content.length} bytes`);
    console.log(`Object content: ${content}`);

    // Verify content
    if (content === objectBody) {
        console.log("Content verification: SUCCESS");
    } else {
        console.log("Content verification: FAILED");
        console.log(`Expected: ${objectBody}`);
        console.log(`Actual: ${content}`);
    }
}

async function deleteObject() {
    const command = new DeleteObjectCommand({
        Bucket: bucketName,
        Key: objectKey
    });
    
    await s3Client.send(command);
    console.log(`Object successfully deleted: ${objectKey}`);
}

main(); 