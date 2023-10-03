package test

import (
	"com.blocopad/blocopad/internal/backend"
	"com.blocopad/blocopad/internal/db"
	"testing"
)

func TestGetKey(t *testing.T) {
	// Given
	db.GetNote = func(key string) (bool, string, error) {
		return false, "OK", nil
	}
	// When
	data, err := backend.GetKey("Key1")

	// Then
	if err != nil {
		t.Fatal("TestGetKeyOk Should nor return error")
	}
	if data != "OK" {
		t.Fatal("TestGetKeyOk Invalid return")
	}
}
