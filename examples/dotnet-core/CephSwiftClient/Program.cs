using System;
using System.Threading.Tasks;
using Amazon.S3;
using Amazon.S3.Model;
using Amazon.Runtime;
using System.IO;

class Program
{
    private static readonly string EndpointUrl = "https://pvos-tr-ist-01.portvmind.com";
    private static readonly string AccessKeyId = "<ACCESS_KEY_ID>";
    private static readonly string SecretKey = "<SECRET_KEY>";
    private static readonly string Region = "tr-ist-01";
    private static readonly string TestBucketName = "test-bucket-dotnet";
    private static readonly string TestFileName = "test.txt";
    private static readonly string TestFileContent = "This is a test file content for Ceph S3 API testing.";

    static async Task Main(string[] args)
    {
        try
        {
            Console.WriteLine("Starting Ceph S3 API Test...");
            var s3Client = CreateS3Client();

            // Test 1: List existing buckets
            Console.WriteLine("\nTest 1: Listing existing buckets");
            await ListBuckets(s3Client);

            // Test 2: Create a test bucket
            Console.WriteLine("\nTest 2: Creating a test bucket");
            await CreateBucket(s3Client, TestBucketName);

            // Test 3: Upload a test file
            Console.WriteLine("\nTest 3: Uploading a test file");
            byte[] content = System.Text.Encoding.UTF8.GetBytes(TestFileContent);
            await UploadObject(s3Client, TestBucketName, TestFileName, content);

            // Test 4: Download and verify the test file
            Console.WriteLine("\nTest 4: Downloading and verifying the test file");
            await DownloadObject(s3Client, TestBucketName, TestFileName);

            // Test 5: Delete the test file
            Console.WriteLine("\nTest 5: Deleting the test file");
            await DeleteObject(s3Client, TestBucketName, TestFileName);

            Console.WriteLine("\nAll tests completed successfully!");
        }
        catch (Exception ex)
        {
            Console.WriteLine($"\nAn error occurred during testing: {ex.Message}");
            Console.WriteLine($"Stack trace: {ex.StackTrace}");
        }
    }

    static IAmazonS3 CreateS3Client()
    {
        var config = new AmazonS3Config
        {
            ServiceURL = EndpointUrl,
            ForcePathStyle = true, // Required for Ceph S3
            UseHttp = true
        };

        var credentials = new BasicAWSCredentials(AccessKeyId, SecretKey);
        return new AmazonS3Client(credentials, config);
    }

    static async Task ListBuckets(IAmazonS3 s3Client)
    {
        var response = await s3Client.ListBucketsAsync();
        Console.WriteLine("Available buckets:");
        foreach (var bucket in response.Buckets)
        {
            Console.WriteLine($"- {bucket.BucketName} (Created: {bucket.CreationDate})");
        }
    }

    static async Task UploadObject(IAmazonS3 s3Client, string bucketName, string objectName, byte[] content)
    {
        var request = new PutObjectRequest
        {
            BucketName = bucketName,
            Key = objectName,
            InputStream = new MemoryStream(content)
        };

        var response = await s3Client.PutObjectAsync(request);
        Console.WriteLine($"File successfully uploaded: {objectName}");
        Console.WriteLine($"ETag: {response.ETag}");
    }

    static async Task DownloadObject(IAmazonS3 s3Client, string bucketName, string objectName)
    {
        var request = new GetObjectRequest
        {
            BucketName = bucketName,
            Key = objectName
        };

        using var response = await s3Client.GetObjectAsync(request);
        using var responseStream = response.ResponseStream;
        using var reader = new StreamReader(responseStream);
        var content = await reader.ReadToEndAsync();

        Console.WriteLine($"File successfully downloaded: {objectName}");
        Console.WriteLine($"File size: {content.Length} bytes");
        Console.WriteLine($"File content: {content}");
        
        // Verify content
        if (content == TestFileContent)
        {
            Console.WriteLine("Content verification: SUCCESS");
        }
        else
        {
            Console.WriteLine("Content verification: FAILED");
            Console.WriteLine($"Expected: {TestFileContent}");
            Console.WriteLine($"Actual: {content}");
        }
    }

    static async Task DeleteObject(IAmazonS3 s3Client, string bucketName, string objectName)
    {
        var request = new DeleteObjectRequest
        {
            BucketName = bucketName,
            Key = objectName
        };

        await s3Client.DeleteObjectAsync(request);
        Console.WriteLine($"File successfully deleted: {objectName}");
    }

    static async Task CreateBucket(IAmazonS3 s3Client, string bucketName)
    {
        try
        {
            var request = new PutBucketRequest
            {
                BucketName = bucketName,
                UseClientRegion = true
            };

            await s3Client.PutBucketAsync(request);
            Console.WriteLine($"Bucket successfully created: {bucketName}");
        }
        catch (AmazonS3Exception ex) when (ex.ErrorCode == "BucketAlreadyOwnedByYou")
        {
            Console.WriteLine($"Bucket already exists: {bucketName}");
        }
    }
}
