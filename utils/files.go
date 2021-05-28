package utils

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func FileListInFolder(root string) (*[]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return &files, err
}

func FileLines(path string) (*[]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return &lines, scanner.Err()
}

func ReadFileToString(path string) (*string, error) {
	content := ""
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += "\n" + scanner.Text()
	}
	return &content, scanner.Err()
}
