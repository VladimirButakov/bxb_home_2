package saver

import (
	"fmt"
	"io"
	"os"
)

type ImageSaver interface {
	Save(hash string, savePath string, content io.Reader) error
}

type FileImageSaver struct{}

func (s *FileImageSaver) Save(hash string, savePath string, content io.Reader) error {
	err := os.MkdirAll(savePath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create image folder: %w", err)
	}
	filePath := fmt.Sprintf("%s/%s.jpg", savePath, hash)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file for image: %w", err)
	}
	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return fmt.Errorf("error saving image: %w", err)
	}

	return nil
}
