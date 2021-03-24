# Convert

This function converts JSON into YAML.

This lambda does the following things in order:
- download the s3 object
- convert JSON to YAML
- write the response back with the correct token and route