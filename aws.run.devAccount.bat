@echo off
aws --profile devAccount --region ap-southeast-2 lambda invoke --function-name go-lambda-bcryptit --payload file://in.json _out.json