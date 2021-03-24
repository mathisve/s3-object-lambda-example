// runExamples.go
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"net/http"
)

// correct input struct
// the actual incoming json is a lot longer, but we only need to use these 3 parameters
type Input struct {
	GetObjectContext struct {
		Inputs3URL  string `json:"inputS3Url"`
		OutputRoute string `json:"outputRoute"`
		OutputToken string `json:"outputToken"`
	} `json:"GetObjectContext"`
}

type Person struct {
		Name string `json:"name"`
		Age int `json:"age"`
		Location string `json:"location"`
}

var s3session *s3.S3

const (
	REGION = "eu-central-1"
)

func init() {
	// create a new s3 session
	s3session = s3.New(session.Must(session.NewSession(&aws.Config{
		Region: aws.String(REGION),
	})))
}

func handler(ctx context.Context, event Input){

	// downloads the requested object
	received, err := GetFile(event.GetObjectContext.Inputs3URL)
	if err != nil {
		log.Println(err)
		return
	}

	// convert json to yaml

	var p []Person
	err = json.Unmarshal(received, &p)
	if err != nil {
		log.Println(err)
		return
	}

	edited, err := yaml.Marshal(p)
	if err != nil {
		log.Println(err)
		return
	}

	// writes the response with the correct route and token
	_, err = s3session.WriteGetObjectResponseWithContext(ctx, &s3.WriteGetObjectResponseInput{
		Body: bytes.NewReader(edited),
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

	// reads the bytes
	b, err = ioutil.ReadAll(resp.Body)
	return b, err
}

func main() {
	lambda.Start(handler)
}