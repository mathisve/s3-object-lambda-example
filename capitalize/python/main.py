import boto3
from botocore.vendored import requests

def lambda_handler(event, context):
    object_get_context = event["getObjectContext"]
    request_route = object_get_context["outputRoute"]
    request_token = object_get_context["outputToken"]
    s3_url = object_get_context["inputS3Url"]

    response = requests.get(s3_url)
    original_object = response.content.decode("utf-8")

    transformed_object = original_object.upper()

    s3 = boto3.client('s3')
    s3.write_get_object_response(
        Body=transformed_object,
        RequestRoute=request_route,
        RequestToken=request_token)

    return {'status_code': 200}