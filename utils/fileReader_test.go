package utils

import "testing"

// Change the folderName and name to test the function
func TestGetVariable(t *testing.T) {
	var folderName string = "config"
	var name string = "test"
	var expected string = "test"
	var actual string
	var err error
	actual, err = GetVariable(folderName, name)
	if err != nil {
		t.Errorf("GetVariable() error: %v", err)
	}
	if actual != expected {
		t.Errorf("GetVariable() actual: %v, expected: %v", actual, expected)
	}
}
