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
	Path string // Exported field
}

type Manager interface {
	SaveBinJson(bin Bin)
	LoadBinsFromJson(P PathStruct) (BinList, error)
}

func NewPathStruct(path string) PathStruct {
	return PathStruct{Path: path}
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
	bytes, err := json.MarshalIndent(data, "", "  ") // Added indentation for readability
	if err != nil {
		fmt.Println(err)
	}
	return bytes
}

func (P PathStruct) LoadBinsFromJson() (BinList, error) {
	var binList BinList
	data, err := os.ReadFile(P.Path)
	if err != nil {
		return binList, err
	}
	err = json.Unmarshal(data, &binList)
	if err != nil {
		return binList, err
	}
	return binList, nil
}

func LoadBinsFromPath(path string) (BinList, error) {
	pathStruct := NewPathStruct(path)
	return pathStruct.LoadBinsFromJson()
}
