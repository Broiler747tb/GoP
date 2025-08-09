package file

import (
	"GoP/bins"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

func CreateUser(id string, private bool, createdAt time.Time, name string) *bins.Bin {
	return &bins.Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}
}

func ReadJsonFile(path string) ([]byte, error) {
	var returnByte []byte
	if strings.HasSuffix(path, ".json") {
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		jsonBytes, err := io.ReadAll(file)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(jsonBytes, &jsonBytes)
		if err != nil {
			return nil, err
		}
		returnByte = jsonBytes
	}
	return returnByte, nil
}
