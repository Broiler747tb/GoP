package file

import (
	"GoP/bins"
	"encoding/json"
	"fmt"
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
	if !strings.HasSuffix(path, ".json") {
		return nil, fmt.Errorf("error file is not .json")
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var temp interface{}
	err = json.Unmarshal(bytes, &temp)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
