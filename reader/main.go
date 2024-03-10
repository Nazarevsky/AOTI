package reader

import (
	"errors"
	"fmt"
	"os"
)

func ReadBinaryFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to open a file: %s\n", err))
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to get file size: %s\n", err))
	}
	fileSize := fileInfo.Size()

	content := make([]byte, fileSize)
	_, err = file.Read(content)
	if err != nil {
		return "", errors.New(fmt.Sprintf("failed to read a file: %s\n", err))
	}

	contentString := string(content)

	return contentString, nil
}
