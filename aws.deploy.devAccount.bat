@echo off
aws --profile devAccount --region ap-southeast-2 lambda update-function-code --function-name go-lambda-bcryptit --zip-file fileb://main.zip --publish