package raw

import (
	"os"
	"path/filepath"
)

func Write(filename string, data []byte) error {
	// Create missing nested directories if needed
	dirname := filepath.Dir(filename)
	err := os.MkdirAll(dirname, os.ModePerm)

	if err != nil {
		return err
	}

	file, err := os.Create(filename)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.Write(data)

	if err != nil {
		return err
	}

	return nil
}
