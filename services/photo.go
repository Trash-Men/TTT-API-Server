package services

import (
	"mime/multipart"
	"strings"
	"time"

	"github.com/Trash-Men/api-server/configs"
	"github.com/gofrs/uuid"
)

const imageDirectoryPath string = "upload/"

var imageService ImageService

func init() {
	if configs.GetEnvironments().MODE == "production" {
		bucketName := configs.GetEnvironments().S3_BUCKET_NAME
		s3Service := S3{}

		s3Service.InitSession()

		error := s3Service.CreateBucket(bucketName)

		if error == nil {
			s3Service.ConfigS3PublicAccess(bucketName)
		}

		imageService = s3Service
	} else {
		imageService = DevelopmentImageService{}
	}
}

type PhotoService struct{}

func (_ PhotoService) UploadPhoto(photo *multipart.FileHeader, photoType string) (string, error) {
	file, error := photo.Open()

	if error != nil {
		return "", error
	}

	defer file.Close()

	uuidResponse, uuidError := uuid.NewV4()

	if uuidError != nil {
		return "", uuidError
	}

	fileExtension := strings.Split(photo.Filename, ".")
	fileName := uuidResponse.String() + "." + fileExtension[len(fileExtension)-1]

	currentDatePath := strings.Join(strings.Split(time.Now().Format("2006-01-02"), "-"), "/")

	photoPath := imageDirectoryPath + photoType + "/" + currentDatePath + "/" + fileName

	if error := imageService.UploadImage(file, photoPath); error != nil {
		return "", error
	}

	return photoPath, nil
}

func (_ PhotoService) DeletePhoto(photoPath string) error {
	return imageService.DeleteImage(photoPath)
}
