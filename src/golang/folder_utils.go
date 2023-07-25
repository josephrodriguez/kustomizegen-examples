package main

import (
	"fmt"
	"os"
)

func createFolder(folderName string, perm os.FileMode) error {
	err := os.Mkdir(folderName, perm)
	if err != nil {
		return err
	}
	return nil
}

func validateFolderExists(folderPath string) error {
	// Check if the folder exists
	_, err := os.Stat(folderPath)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("folder does not exist: %s", folderPath)
		}
		return err
	}
	return nil
}
