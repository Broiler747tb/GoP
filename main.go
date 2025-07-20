package main

import "time"

func main() {

}

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func createUser(id string, private bool, createdAt time.Time, name string) *Bin {
	return &Bin{
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}
}
