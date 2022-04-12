package util

import (
	_ "crypto/sha256"
	_ "crypto/sha512"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/opencontainers/go-digest"
)

type FileDetails struct {
	Digest   digest.Digest
	Size     int64
	MimeType string
}

func GetDigest(data []byte) digest.Digest {
	return digest.FromBytes(data)
}

func GetMimeType(data []byte) string {
	return mimetype.Detect(data).String()
}

func GetSize(data []byte) int64 {
	return int64(len(data))
}

func GetFileDetails(filePath string) (*FileDetails, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return GetByteDetails(data), nil
}

func GetByteDetails(data []byte) *FileDetails {
	details := FileDetails{
		Digest:   GetDigest(data),
		Size:     GetSize(data),
		MimeType: GetMimeType(data),
	}
	return &details
}
