package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"strings"

	"github.com/Trash-Men/api-server/configs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type S3 struct {
	session *session.Session
	client  *s3.S3
}

func (_s3 *S3) InitSession() {
	environments := configs.GetEnvironments()

	_session, error := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials(environments.IAM_ACCESS_KEY, environments.IAM_SECRET_ACCESS_KEY, ""),
	})

	if error != nil {
		os.Exit(1)
	}

	_s3.session = _session
	_s3.client = s3.New(_session)
}

func (_s3 S3) CreateBucket(bucketName string) error {
	_, error := _s3.client.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(bucketName),
	})

	if error != nil {
		matchedString := "status code:"

		startIndex := strings.Index(error.Error(), matchedString) + len(matchedString) + 1
		STATUS_CODE_LENGTH := 3

		statusCode := error.Error()[startIndex : startIndex+STATUS_CODE_LENGTH]

		if statusCode == "409" {
			return errors.New("already existed")
		}

		panic(error.Error())
	}

	error = _s3.client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucketName),
	})

	if error != nil {
		panic(error.Error())
	}

	return nil
}

func (_s3 S3) ConfigS3PublicAccess(bucketName string) {
	readOnlyAnonUserPolicy := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Sid":       "AddPerm",
				"Effect":    "Allow",
				"Principal": "*",
				"Action": []string{
					"s3:GetObject",
				},
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s/*", bucketName),
				},
			},
		},
	}

	policy, err := json.Marshal(readOnlyAnonUserPolicy)

	if err != nil {
		panic(err.Error())
	}

	_, err = _s3.client.PutBucketPolicy(&s3.PutBucketPolicyInput{
		Bucket: aws.String(bucketName),
		Policy: aws.String(string(policy)),
	})

	fmt.Printf("Successfully set bucket %q's policy\n", bucketName)
}

func (_s3 S3) UploadImage(file multipart.File, photoPath string) error {
	uploader := s3manager.NewUploader(_s3.session)
	bucket := configs.GetEnvironments().S3_BUCKET_NAME

	_, error := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(photoPath),
		Body:   file,
	})

	if error != nil {
		return error
	}

	fmt.Printf("Successfully uploaded %q to %q\n", photoPath, bucket)
	return nil
}

func (_s3 S3) DeleteImage(imagePath string) error {
	bucketName := configs.GetEnvironments().S3_BUCKET_NAME

	response, error := _s3.client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucketName)})
	if error != nil {
		return error
	}

	isExist := false

	for _, item := range response.Contents {
		if *item.Key == imagePath {
			isExist = true
		}
	}

	if !isExist {
		return fmt.Errorf("%s not existed", imagePath)
	}

	_, error = _s3.client.DeleteObject(&s3.DeleteObjectInput{Bucket: aws.String(bucketName), Key: aws.String(imagePath)})

	if error != nil {
		return error
	}

	error = _s3.client.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imagePath),
	})

	return error
}
