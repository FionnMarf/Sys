package main

import (
	"fmt"
	"os"
)

func walkfunction(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()
	if mode.IsDir() {
		fmt.Println(path)
	}
	return nil
}
