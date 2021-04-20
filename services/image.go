package services

import "mime/multipart"

type ImageService interface {
	UploadImage(file multipart.File, photoPath string) error
	DeleteImage(imatPath string) error
}
