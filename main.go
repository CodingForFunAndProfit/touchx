package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Printf(`Usage: %s filepath/filename
This program only accepts one parameter. 
The text after the last slash is the file that is created.
This program does not use the shell command mkdir or touch, yet.`, filepath.Base(os.Args[0]))
		os.Exit(0)
	}

	filewithpath := os.Args[1]

	if FileExist(filewithpath) {
		if IsDir(filewithpath) {
			log.Fatalln(filewithpath, "is a directory")
		}
		log.Fatalln("file", filewithpath, "already exists.")
	}
	fileslice := strings.Split(filewithpath, "/")

	if len(fileslice) > 1 { // there is a path
		for i, dirName := range fileslice {
			if i < len(fileslice)-1 { //last slice is the filename
				if !DirExist(dirName) {
					err := os.Mkdir(dirName, os.ModeDir)
					if err != nil {
						log.Fatal(err)
					}
				}
			}
		}
	}
	// attention. truncates an already existing file
	file, err := os.Create(filewithpath)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(filewithpath, "created.")
	file.Close()
}

func FileExist(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DirExist(name string) bool {
	_, err := os.Stat(name)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func IsDir(name string) bool {
	info, err := os.Stat(name)
	if os.IsExist(err) && !info.IsDir() {
		return false
	}
	return true
}
