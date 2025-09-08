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

const BINS_FILE = "bins.json"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// generateID creates a simple random ID
func generateID() string {
	return strconv.Itoa(rand.Intn(1000000))
}

func Create(filen *string, name *string) {
	if filen == nil || *filen == "" {
		fmt.Println("Error: file parameter is required")
		return
	}
	if name == nil || *name == "" {
		fmt.Println("Error: name parameter is required")
		return
	}

	// Read the file content
	fileContent, err := file.ReadJsonFile(*filen)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", *filen, err)
		return
	}

	// Create new bin
	newBin := storage.Bin{
		Bin: bins.Bin{
			Id:        generateID(),
			Private:   false,
			CreatedAt: time.Now(),
			Name:      *name,
		},
	}

	// Load existing bins
	binList := loadOrCreateBinsList()

	// Add new bin to the list
	binList.Bins = append(binList.Bins, newBin.Bin)

	// Save the updated bin list
	if err := saveBinsList(binList); err != nil {
		fmt.Printf("Error saving bin list: %v\n", err)
		return
	}

	// Save the bin content to a separate file
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

	// Load existing bins
	binList := loadOrCreateBinsList()

	// Find the bin to update
	found := false
	for i, bin := range binList.Bins {
		if bin.Id == *id {
			// Read new file content
			fileContent, err := file.ReadJsonFile(*filen)
			if err != nil {
				fmt.Printf("Error reading file %s: %v\n", *filen, err)
				return
			}

			// Update the bin's timestamp
			binList.Bins[i].CreatedAt = time.Now()

			// Save updated bin list
			if err := saveBinsList(binList); err != nil {
				fmt.Printf("Error saving bin list: %v\n", err)
				return
			}

			// Update the bin content file
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

	// Load existing bins
	binList := loadOrCreateBinsList()

	// Find and remove the bin
	found := false
	for i, bin := range binList.Bins {
		if bin.Id == *id {
			// Remove bin from slice
			binList.Bins = append(binList.Bins[:i], binList.Bins[i+1:]...)

			// Save updated bin list
			if err := saveBinsList(binList); err != nil {
				fmt.Printf("Error saving bin list: %v\n", err)
				return
			}

			// Delete the bin content file
			binFileName := fmt.Sprintf("bin_%s.json", *id)
			if err := os.Remove(binFileName); err != nil {
				// Don't fail if file doesn't exist
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

	// Load existing bins
	binList := loadOrCreateBinsList()

	// Find the bin
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

	// Display bin metadata
	fmt.Printf("Bin ID: %s\n", targetBin.Id)
	fmt.Printf("Name: %s\n", targetBin.Name)
	fmt.Printf("Private: %v\n", targetBin.Private)
	fmt.Printf("Created At: %s\n", targetBin.CreatedAt.Format(time.RFC3339))

	// Try to read and display bin content
	binFileName := fmt.Sprintf("bin_%s.json", *id)
	content, err := os.ReadFile(binFileName)
	if err != nil {
		fmt.Printf("Warning: Could not read bin content: %v\n", err)
		return
	}

	fmt.Printf("Content:\n%s\n", string(content))
}

func List() {
	// Load existing bins
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

// Helper function to load or create bins list
func loadOrCreateBinsList() storage.BinList {
	binList := storage.BinList{}

	if _, err := os.Stat(BINS_FILE); os.IsNotExist(err) {
		// File doesn't exist, return empty list
		return binList
	}

	data, err := os.ReadFile(BINS_FILE)
	if err != nil {
		log.Printf("Error reading bins file: %v", err)
		return binList
	}

	if err := json.Unmarshal(data, &binList); err != nil {
		log.Printf("Error unmarshaling bins: %v", err)
		return storage.BinList{} // Return empty list on error
	}

	return binList
}

// Helper function to save bins list
func saveBinsList(binList storage.BinList) error {
	data, err := json.MarshalIndent(binList, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling bins: %v", err)
	}

	if err := os.WriteFile(BINS_FILE, data, 0644); err != nil {
		return fmt.Errorf("error writing bins file: %v", err)
	}

	return nil
}

// Api function to initialize config (call this first)
func Api() {
	config.GetConfig()
	fmt.Println("API initialized successfully")
}
