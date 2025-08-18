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

type Bin struct {
	bins.Bin
}

type PathStruct struct {
	path string
}

type Manager interface {
	SaveBinJson(bin Bin)
	LoadBinsFromJson(filename string) (BinList, error)
}

func SaveBinJson(bin Bin) {
	file, err := os.Create("Bin.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	bytes := ToJson(bin)
	file.Write(bytes)
}

func ToJson(data any) []byte {
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func (P PathStruct) LoadBinsFromJson() (BinList, error) {
	var binList BinList
	data, err := os.ReadFile(P.path)
	if err != nil {
		return binList, err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList, err
	}
	return binList, nil
}
