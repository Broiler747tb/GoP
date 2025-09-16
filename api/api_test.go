package api

import "testing"

func TestCreate(t *testing.T) {
	filen := "user.json"
	name := "john"
	err := Create(&filen, &name)
	if err != nil {
		t.Errorf("There was an error creating a file")
	}
}

func TestUpdate(t *testing.T) {
	filen := "bin_671203.json"
	id := "671203"
	err := Update(&filen, &id)
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
