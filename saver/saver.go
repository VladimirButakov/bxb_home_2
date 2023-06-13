package saver

import (
	"fmt"
	"io"
	"os"
)

type ImageSaver interface {
	Save(filePath string, content io.Reader) error
}

type FileImageSaver struct{}

func (s *FileImageSaver) Save(filePath string, content io.Reader) error {
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
