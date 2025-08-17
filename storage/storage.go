package storage

import (
	"GoP/bins"
	"encoding/json"
	"fmt"
	"os"
)

type BinList struct {
	Bins []bins.Bin `json:"bins"`
}

type BinInterface interface {
}

func SaveBinJson(Bin bins.Bin) {
	file, err := os.Create("Bin.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bytes := ToJson(Bin)
	file.Write(bytes)
}

func ToJson(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func LoadBinsFromJson(filename string) (BinList, error) {
	var binList BinList
	data, err := os.ReadFile(filename)
	if err != nil {
		return binList, err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList, err
	}
	return binList, nil
}
