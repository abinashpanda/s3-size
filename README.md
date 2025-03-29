# s3-size

A command-line utility to analyze and report the size of S3 buckets and their objects.

## Overview

s3-size provides a simple way to measure the storage consumption of your S3 buckets, directories, and individual objects.

## Usage

```sh
NAME:
   s3-size - cli tool to get the size of the s3 bucket and its objects (or directory)

USAGE:
   s3-size [global options] command [command options]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --endpoint value    s3 endpoint
   --access-key value  s3 access key
   --secret-key value  s3 secret key
   --bucket value      root bucket
   --use-ssl           use HTTPS connection when communicating with S3 (default: false)
   --max-depth value   print the size for a directory (or file, with --all) only if it is N or fewer levels below the command line argument (default: 0)
   --help, -h          show help
Required flags "endpoint, access-key, secret-key, bucket" not set
```

## Installation

### Building from Source

You can build the CLI tool locally:

```sh
go build
```

This will generate an `s3-size` executable binary in your current directory.

### Prerequisites

- Go 1.x or higher
- Access to an S3-compatible storage service

## Features

- Calculate total bucket size
- Analyze directory and object sizes
- Support for S3-compatible storage services
- Configurable depth for detailed analysis
- Secure connection option via SSL/TLS

## Examples

Calculate the size of an entire bucket:

```sh
./s3-size --endpoint s3.amazonaws.com --access-key YOUR_ACCESS_KEY --secret-key YOUR_SECRET_KEY --bucket your-bucket-name
```

Analyze with SSL enabled:

```sh
./s3-size --endpoint s3.amazonaws.com --access-key YOUR_ACCESS_KEY --secret-key YOUR_SECRET_KEY --bucket your-bucket-name --use-ssl
```

Set maximum depth for directory analysis:

```sh
./s3-size --endpoint s3.amazonaws.com --access-key YOUR_ACCESS_KEY --secret-key YOUR_SECRET_KEY --bucket your-bucket-name --max-depth 2
```
