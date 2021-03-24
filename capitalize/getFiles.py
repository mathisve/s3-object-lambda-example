import boto3
import time

s3 = boto3.client("s3")

print("Original object from the S3 bucket:")
startTime = time.time()
original = s3.get_object(
    Bucket="s3testbucketmathis",
    Key="fileInS3.txt")
print(original["Body"].read().decode("utf-8"))
print("Took: {}".format(time.time() - startTime))

print("/n")

print("Object processed by S3 Object Lambda:")
startTime = time.time()
transformed = s3.get_object(
    Bucket="arn:aws:s3-object-lambda:eu-central-1:043039367084:accesspoint/lambdaap",
    Key="fileInS3.txt")
print(transformed["Body"].read().decode("utf-8"))
print("Took: {}".format(time.time() - startTime))