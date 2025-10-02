package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	for {
		FilePath, _ := os.Getwd()
		FilesForPath, _ := os.ReadDir(FilePath)
		for index, file := range FilesForPath {
			fileinfo, err := os.Stat(file.Name())
			if err != nil {
				fmt.Println("Error:", err)
				return
			}

			FileExtension := filepath.Ext(file.Name())
			IsDirectory := fileinfo.IsDir()

			if !IsDirectory {
				fmt.Println(index, "File name: ", fileinfo.Name(), "Mod time: ", fileinfo.ModTime(), "File size: ", fileinfo.Size(), "Mode: ", fileinfo.Mode(), "File extension:", FileExtension)
			} else if IsDirectory {
				fmt.Println(index, "File name: ", fileinfo.Name(), "Mod time: ", fileinfo.ModTime(), "File size: ", fileinfo.Size(), "Mode: ", fileinfo.Mode(), "File extension:", "directory")
			} else {
				fmt.Println("Error")
			}

		}
		r := bufio.NewReader(os.Stdin)
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		CommandHandler(line)
	}
}
