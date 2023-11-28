package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	err := ListFileDirectory("/home/rafael/Documentos/aulajava")

	if err != nil {
		fmt.Println(err)
	}

}

func ListFileDirectory(directory string) error {
	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
