package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws"
	"os"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var LOG = fmt.Println

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Region: aws.String("us-west-2"),
		},
		Profile:           "default",
		SharedConfigState: session.SharedConfigEnable,
	}))

	credential, err := sess.Config.Credentials.Get()

	if err != nil {
		panic(err)
	}

	LOG(credential)

	s3Service := s3.New(sess, &aws.Config{
		Region: aws.String("us-west-2"),
	})

	result, err := s3Service.ListBuckets(nil)

	if err != nil {
		panic(err)
	}

	bucket := result.Buckets[0]

	filename := "names.txt"
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: bucket.Name,
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

}
