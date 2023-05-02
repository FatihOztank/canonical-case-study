package main

import (
	"example/exercise2/shred"
	"testing"
)

// this file tests the basic functionality of the shred function
// run "go test" to run tests.
func TestMain(t *testing.T) {
	testFile := "./test.txt"
	err := shred.CopyFile("input.txt", testFile)
	if err != nil {
		t.Error("Copy failed...")
	}
	err = shred.Shred(testFile)
	if err != nil {
		t.Error("Shred failed...")
	}
}
