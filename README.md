# go lambda bcryptit

An AWS lambda function that sits behind the API Gateway and performs the following:
- Given a password, return its bcrypt hash
- Given a password and a hash, verify that they are valid

# Usage

| Path | Expected input | Expected output | Description
|---|---|---|---|
| /bcrypt/hash | ``` { "password" : "secret" } ```| ``` {"result" : "@#$@#$" } ```| Returns a bcrypt hash of the provided password |
| /bcrypt/verify | ``` { "password" : "secret", "hash" : "$#$%#$%" } ```| ``` {"message" : "valid" } ``` or ``` {"message" : "invalid" } ```| Returns a bcrypt hash of the provided password |


The API GW should be configured as follows:
```
/
  /bcrypt
  OPTIONS
    /{proxy+}
    ANY
    OPTIONS
```
where ANY method is integrated with our Lambda via lambda proxy integration


# Deployment to AWS

## Simply upload
- Create a new lambda function with go1.x as runtime
- Upload main.zip
- Rename "handler" to "main"
- And your lambda is ready to be used

If you fancy, use the deploy script :)

## Build from source
- Run the build.cmd script to create an executable
- The same script should work on a linux system as well
- Then deploy

# License
Refer to LICENSE file