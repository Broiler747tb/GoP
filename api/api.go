package api

import (
	"GoP/bins"
	"GoP/config"
	"GoP/file"
	"GoP/storage"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const binsFile = "bins.json"

func generateID() string {
	return strconv.Itoa(rand.Intn(1000000))
}

func Create(filen *string, name *string) {
	if *filen == "" {
		fmt.Println("Error: file name is required")
		return
	}
	if *name == "" {
		fmt.Println("Error: name is required")
		return
	}
	fileContent, err := file.ReadJsonFile(*filen)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *filen, err)
		return
	}
	newBin := storage.Bin{
		Bin: bins.Bin{
			Id:        generateID(),
			Private:   false,
			CreatedAt: time.Now(),
			Name:      *name,
		},
	}

	binList := loadOrCreateBinsList()
	binList.Bins = append(binList.Bins, newBin.Bin)

	if err := saveBinsList(binList); err != nil {
		fmt.Printf("Error saving bin list: %v\n", err)
		return
	}

	binFileName := fmt.Sprintf("bin_%s.json", newBin.Id)
	if err := os.WriteFile(binFileName, fileContent, 0644); err != nil {
		fmt.Printf("Error saving bin content: %v\n", err)
		return
	}

	fmt.Printf("Created bin with ID: %s, Name: %s\n", newBin.Id, newBin.Name)
}

func Update(filen *string, id *string) {
	if filen == nil || *filen == "" {
		fmt.Println("Error: file parameter is required")
		return
	}
	if id == nil || *id == "" {
		fmt.Println("Error: id parameter is required")
		return
	}

	binList := loadOrCreateBinsList()

	found := false
	for i, bin := range binList.Bins {
		if bin.Id == *id {

			fileContent, err := file.ReadJsonFile(*filen)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", *filen, err)
				return
			}

			binList.Bins[i].CreatedAt = time.Now()

			if err := saveBinsList(binList); err != nil {
				fmt.Printf("Error saving bin list: %v\n", err)
				return
			}

			binFileName := fmt.Sprintf("bin_%s.json", *id)
			if err := os.WriteFile(binFileName, fileContent, 0644); err != nil {
				fmt.Printf("Error updating bin content: %v\n", err)
				return
			}

			fmt.Printf("Updated bin with ID: %s\n", *id)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Bin with ID %s not found\n", *id)
	}
}

func Delete(id *string) {
	if id == nil || *id == "" {
		fmt.Println("Error: id parameter is required")
		return
	}

	binList := loadOrCreateBinsList()

	found := false
	for i, bin := range binList.Bins {
		if bin.Id == *id {
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)

			if err := saveBinsList(binList); err != nil {
				fmt.Printf("Error saving bin list: %v\n", err)
				return
			}

			binFileName := fmt.Sprintf("bin_%s.json", *id)
			if err := os.Remove(binFileName); err != nil {

				if !os.IsNotExist(err) {
					fmt.Printf("Warning: Error deleting bin file: %v\n", err)
				}
			}

			fmt.Printf("Deleted bin with ID: %s\n", *id)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Bin with ID %s not found\n", *id)
	}
}

func Get(id *string) {
	if id == nil || *id == "" {
		fmt.Println("Error: id parameter is required")
		return
	}

	binList := loadOrCreateBinsList()

	var targetBin *bins.Bin
	for _, bin := range binList.Bins {
		if bin.Id == *id {
			targetBin = &bin
			break
		}
	}

	if targetBin == nil {
		fmt.Printf("Bin with ID %s not found\n", *id)
		return
	}

	fmt.Printf("Bin ID: %s\n", targetBin.Id)
	fmt.Printf("Name: %s\n", targetBin.Name)
	fmt.Printf("Private: %v\n", targetBin.Private)
	fmt.Printf("Created At: %s\n", targetBin.CreatedAt.Format(time.RFC3339))

	binFileName := fmt.Sprintf("bin_%s.json", *id)
	content, err := os.ReadFile(binFileName)
	if err != nil {
		fmt.Printf("Warning: Could not read bin content: %v\n", err)
		return
	}

	fmt.Printf("Content:\n%s\n", string(content))
}

func List() {
	binList := loadOrCreateBinsList()

	if len(binList.Bins) == 0 {
		fmt.Println("No bins found")
		return
	}

	fmt.Printf("Found %d bin(s):\n", len(binList.Bins))
	fmt.Println("ID\t\tName\t\tPrivate\tCreated At")
	fmt.Println("--\t\t----\t\t-------\t----------")

	for _, bin := range binList.Bins {
		fmt.Printf("%s\t\t%s\t\t%v\t%s\n",
			bin.Id,
			bin.Name,
			bin.Private,
			bin.CreatedAt.Format("2006-01-02 15:04:05"))
	}
}

func loadOrCreateBinsList() storage.BinList {
	binList := storage.BinList{}

	_, err := os.Stat(binsFile)
	if os.IsNotExist(err) {
		return binList
	}

	data, err := os.ReadFile(binsFile)
	if err != nil {
		log.Printf("Error reading bins file: %v", err)
		return binList
	}

	err = json.Unmarshal(data, &binList)
	if err != nil {
		log.Printf("Error unmarshaling bins: %v", err)
		return storage.BinList{}
	}

	return binList
}

func saveBinsList(binList storage.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling bins: %v", err)
	}

	if err := os.WriteFile(binsFile, data, 0644); err != nil {
		return fmt.Errorf("error writing bins file: %v", err)
	}

	return nil
}

func Api() {
	config.GetConfig()
	fmt.Println("API initialized successfully")
}
