package api

import (
	"GoP/bins"
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestCreate(t *testing.T) {
	filen := "user.json"
	name := "john"
	err := Create(&filen, &name)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(filen)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	byteFile, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	binLocal := bins.Bin{}
	err = json.Unmarshal(byteFile, &binLocal)
	if err != nil {
		t.Error(err)
	}
	id := binLocal.Id
	Delete(&id, &filen)
}

func TestUpdate(t *testing.T) {
	filen := "user.json"
	name := "john"
	err := Create(&filen, &name)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(filen)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	byteFile, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	binLocal := bins.Bin{}
	err = json.Unmarshal(byteFile, &binLocal)
	if err != nil {
		t.Error(err)
	}
	id := binLocal.Id
	err = Update(&filen, &id)
	if err != nil {
		t.Error(err)
	}
}

func TestList(t *testing.T) {
	err := List()
	if err != nil {
		t.Error(err)
	}
}

func TestDelete(t *testing.T) {
	filen := "user.json"
	name := "john"
	err := Create(&filen, &name)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(filen)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	byteFile, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	binLocal := bins.Bin{}
	err = json.Unmarshal(byteFile, &binLocal)
	if err != nil {
		t.Error(err)
	}
	id := binLocal.Id
	Delete(&id, &filen)
}

func TestGet(t *testing.T) {
	filen := "user.json"
	name := "john"
	err := Create(&filen, &name)
	if err != nil {
		t.Error(err)
	}
	file, err := os.Open(filen)
	if err != nil {
		t.Error(err)
	}
	defer file.Close()
	byteFile, err := io.ReadAll(file)
	if err != nil {
		t.Error(err)
	}
	binLocal := bins.Bin{}
	err = json.Unmarshal(byteFile, &binLocal)
	if err != nil {
		t.Error(err)
	}
	id := binLocal.Id
	Get(&id)
	Delete(&id, &filen)
}
