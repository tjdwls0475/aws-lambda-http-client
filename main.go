package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
	urlPath := "http://harryhwang.tistory.com/"
	if _, err := url.ParseRequestURI(urlPath); err != nil {
		fmt.Printf("Couldn't parse url. URL is in invalid form.")
		os.Exit(1)
	}
	response, err := http.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		fmt.Print("Status Code not 200")
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%s", body), nil
}

func main() {
	lambda.Start(HandleRequest)
}
