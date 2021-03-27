package utils

import "os"

func DeleteFile(fileName string) error {
	// Removing file
	// Using Remove() function
	err := os.Remove(fileName)
	if err != nil {
		return err
	}
	return nil
}
