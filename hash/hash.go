package hash

import (
	"crypto/md5"
	"fmt"
	"io"
)

type HashCalculator interface {
	CalculateHash(content io.Reader) (string, error)
}

type MD5HashCalculator struct{}

func (h *MD5HashCalculator) CalculateHash(content io.Reader) (string, error) {
	hash := md5.New()
	_, err := io.Copy(hash, content)
	if err != nil {
		return "", fmt.Errorf("error calculating hash: %w", err)
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
