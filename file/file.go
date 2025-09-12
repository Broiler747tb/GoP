package file

import (
	"encoding/json"
	"io"
	"os"
	"strings"
)

type Manager interface {
	ReadJsonFile(path string) ([]byte, error)
}

func ReadJsonFile(path string) ([]byte, error) {
	var returnByte []byte
	if strings.HasSuffix(path, ".json") {
		file, err := os.Open(path)
		defer file.Close()
		if err != nil {
			return nil, err
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		var jsonString string
		err = json.Unmarshal(bytes, &jsonString)
		if err != nil {
			return nil, err
		}
		jsonBytes := []byte(jsonString)
		returnByte = jsonBytes
	} else {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		bytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		returnByte = bytes
	}
	return returnByte, nil
}
