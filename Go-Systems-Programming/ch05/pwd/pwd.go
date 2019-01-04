package main

import (
	"os"
	"fmt"
	"path/filepath"
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	args := os.Args
	pwd, err := os.Getwd()
	checkErr(err)
	if len(args) == 1 {
		fmt.Println(pwd)
		return
	} else {
		if os.Args[1] != "-P" {
			fmt.Println("pwd: bad option: ", os.Args[1])
			os.Exit(1)
		}
		fileinfo, err := os.Lstat(pwd)
		checkErr(err)
		if fileinfo.Mode()&os.ModeSymlink != 0 {
			realpath, err := filepath.EvalSymlinks(pwd)
			checkErr(err)
			fmt.Println(realpath)
		} else {
			fmt.Println(pwd)
		}
	}
}
