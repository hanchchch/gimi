package aws

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/hanchchch/gimi/packages/inspection/pkg/urls"
)

type S3Client struct {
	region string
	sess   *session.Session
}

type S3ClientOptions struct {
	AccessKeyId     string
	SecretAccessKey string
	Region          string
}

func NewS3Client(options S3ClientOptions) (S3Client, error) {
	static_credentials := credentials.NewStaticCredentials(
		options.AccessKeyId,
		options.SecretAccessKey,
		"",
	)

	sess, err := session.NewSession(&aws.Config{
		Credentials: static_credentials,
		Region:      aws.String(options.Region),
	})

	if err != nil {
		return S3Client{}, fmt.Errorf("unable to start aws session: %v", err)
	}

	return S3Client{
		sess:   sess,
		region: options.Region,
	}, nil
}

func (s *S3Client) Upload(bucket string, key string, reader io.Reader) error {
	u := s3manager.NewUploader(s.sess)

	_, err := u.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   reader,
	})

	if err != nil {
		return fmt.Errorf("unable to upload %v", err)
	}

	return nil
}

func (s *S3Client) UploadScreenshot(bucket string, url string, data []byte) (string, error) {
	key := fmt.Sprintf("%v.png", urls.TrimProtocol(url))
	err := s.Upload(
		bucket,
		key,
		ioutil.NopCloser(bytes.NewReader(data)),
	)
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.region, bucket, key), err
}
