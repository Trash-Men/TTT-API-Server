package services

import (
	"errors"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type DevelopmentImageService struct{}

func isNotExist(path string) bool {
	if _, error := os.Stat(path); os.IsNotExist(error) {
		return true
	}

	return false
}

func (developmentImageService DevelopmentImageService) UploadImage(file multipart.File, photoPath string) error {
	direcotries := strings.Split(photoPath, "/")

	for index := range direcotries {
		if index == len(direcotries)-1 {
			break
		}

		if directoryPath := strings.Join(direcotries[0:index+1], "/"); isNotExist(directoryPath) {
			if error := os.Mkdir(directoryPath, 0770); error != nil {
				return error
			}
		}
	}

	destination, error := os.Create(photoPath)

	if error != nil {
		return error
	}

	defer destination.Close()

	if _, error = io.Copy(destination, file); error != nil {
		return error
	}

	file, openFileError := os.Open(photoPath)

	if openFileError != nil {
		return openFileError
	}

	defer file.Close()

	return nil
}

func (developmentImageService DevelopmentImageService) DeleteImage(imagePath string) error {
	if isNotExist(imagePath) {
		return errors.New("not found")
	}

	if error := os.Remove(imagePath); error != nil {
		return errors.New("failed to remove")
	}

	return nil
}
