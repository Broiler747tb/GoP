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
