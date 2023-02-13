# aws-lambda-http-client
Make a http call by using AWS Lambda for free!

# Windows OS Script
$env:GOOS = "linux"
$env:GOARCH = "amd64"
$env:CGO_ENABLED = "0"
go build -o main main.go
~\Go\Bin\build-lambda-zip.exe -o main.zip main

If you don't have build-lambda-zip.exe, then make it by the command: go install github.com/aws/aws-lambda-go/cmd/build-lambda-zip@latest
