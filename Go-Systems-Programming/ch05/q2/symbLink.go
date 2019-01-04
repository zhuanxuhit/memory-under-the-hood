package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		fmt.Println("Please provide an argument!")
		os.Exit(1)
	}
	fileName := args[1]
	fileInfo, err := os.Lstat(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if fileInfo.Mode()&os.ModeSymlink != 0 {
		fmt.Println(fileName, " is a symbolic link.")
		realPath, err := filepath.EvalSymlinks(fileName)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("real path: ", realPath)
	}
}
