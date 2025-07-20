package file

import (
	"GoP/bins"
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
