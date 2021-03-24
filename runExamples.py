import boto3
import time

s3 = boto3.client("s3")

def getObject(bucket, key):
    startTime = time.time()
    original = s3.get_object(
        Bucket=bucket,
        Key=key)
    print(original["Body"].read().decode("utf-8"))
    print("took: {}ms".format(time.time() - startTime))
    print("\n")


print("Original:")
getObject("s3testbucketmathis", "fileInS3.txt")

print("Modified:")
getObject("arn:aws:s3-object-lambda:eu-central-1:043039367084:accesspoint/capitalizeap", "fileInS3.txt")

print("Original:")
getObject("s3testbucketmathis", "people.json")

print("Modified:")
getObject("arn:aws:s3-object-lambda:eu-central-1:043039367084:accesspoint/convertap", "people.json")

