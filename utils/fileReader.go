package utils

import (
	"bufio"
	"os"
)

// Get config variable from file created by configmap
func GetVariable(folderName string, name string) (string, error) {
	file, err := os.Open(folderName + "/" + name)
	if err != nil {
		return "", err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", scanner.Err()
}
