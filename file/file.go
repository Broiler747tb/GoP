package file

import (
	"GoP/bins"
	"encoding/json"
	"io"
	"os"
	"strings"
	"time"
)

type Manager interface {
	CreateUser(id string, private bool, createdAt time.Time, name string) *LocBin
	ReadJsonFile(path string) ([]byte, error)
}

type LocBin struct {
	Bin bins.Bin
}

func (l LocBin) CreateUser(id string, private bool, createdAt time.Time, name string) *LocBin {
	return &LocBin{
		Bin: bins.Bin{
			Id:        id,
			Private:   private,
			CreatedAt: createdAt,
			Name:      name,
		},
	}
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
