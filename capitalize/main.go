// main.go
package main

import (
	"bytes"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Input struct {
	GetObjectContext struct {
		Inputs3URL  string `json:"inputS3Url"`
		OutputRoute string `json:"outputRoute"`
		OutputToken string `json:"outputToken"`
	} `json:"GetObjectContext"`
}

var s3session *s3.S3

const (
	REGION = "eu-central-1"
)

func init() {
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func handler(ctx context.Context, event Input){

	received, err := GetFile(event.GetObjectContext.Inputs3URL)
	if err != nil {
		log.Println(err)
	}

	edited := strings.ToUpper(string(received))
	log.Println(edited)

	_, err = s3session.WriteGetObjectResponseWithContext(ctx, &s3.WriteGetObjectResponseInput{
		Body: bytes.NewReader([]byte(edited)),
		RequestRoute: &event.GetObjectContext.OutputRoute,
		RequestToken: &event.GetObjectContext.OutputToken,
	})

	if err != nil {
		log.Println(err)
	}

}

func GetFile(url string) (b []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return b, err
	}
	defer resp.Body.Close()

	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}

func main() {
	lambda.Start(handler)
}