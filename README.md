* S3 Utility

This small utility allows to upload and download files to/from bucket

## Authentication

### With AWS profile

If a profile can be used (ie `~/.aws/config`), expoert the profile to be used.
```
export AWS_PROFILE=my-profile
```

### With Envvars

The authentication can be passed as EnvVars as following
```
export AWS_ACCESS_KEY_ID=XXX
export AWS_SECRET_ACCESS_KEY=XXX
export AWS_REGION=us-east-1
```

## Usage

### Upload a file

```
s3-utility upload [bucket_name] [local_file]
```

### Download a file

```
s3-utility download [bucket_name] [local_file]
```

It will fail if `local_file` already exists locally

## Install

```
go get gitlab.magicleap.io/ml-compute/s3-utility
```

