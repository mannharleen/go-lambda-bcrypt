@echo off
aws --profile devAccount --region ap-southeast-2 lambda invoke --function-name go-lambda-bcryptit --payload " { \"password\": \"secret\"} " _out.json