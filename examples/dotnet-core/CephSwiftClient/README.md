# Ceph S3 API .NET Core Client

This project is a .NET Core application that performs basic operations using Ceph Object Storage's S3 API.

## Features

- S3 API Authentication
- Bucket listing
- File upload
- File download

## Installation

1. Clone the project
2. Update the following variables in Program.cs according to your Ceph installation:
   - `EndpointUrl`: Ceph S3 endpoint URL
   - `AccessKeyId`: Your S3 access key ID
   - `SecretKey`: Your S3 secret key
   - `Region`: Your S3 region (default: "default")

## Usage

To build and run the project:

```bash
dotnet build
dotnet run
```

## Example Usage

```csharp
// File upload example
byte[] content = File.ReadAllBytes("example.txt");
await UploadObject("bucket-name", "example.txt", content);

// File download example
await DownloadObject("bucket-name", "example.txt");
```

## Requirements

- .NET Core 9.0 or higher
- Ceph Object Storage installation with S3 API enabled
- S3 API credentials (access key and secret key) 