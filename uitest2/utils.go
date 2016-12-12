package main

import (
	"os"
	"path/filepath"
)

func getCurrentDir() string {
	path, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return path
}
