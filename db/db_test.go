package db

import (
	"testing"
)

type Book struct {
	Name string
}

func TestInsert(t *testing.T) {
	book := Book{Name: "go"}
	if err := Insert("test", book); err != nil {
		t.Error("Insert error:", err)
	} else {
		t.Log("Insert Ok")
	}
}
