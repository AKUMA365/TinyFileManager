package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CommandHandler(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}

	switch parts[0] {
	case "cd":
		if len(parts) < 2 {
			fmt.Println("Error: Command \"cd\" requires at least 2 arguments")
			return
		}
		GetCD(parts[1])
	case "pwd":
		GetPwd()
	case "cp":
		if len(parts) < 3 {
			fmt.Println("Error: Command \"cp\" requires at least 3 arguments")
			return
		}
		GetCP(parts[1], parts[2])

	}

}

func GetCD(DirectoryPath string) {
	info, err := os.Stat(DirectoryPath)
	if os.IsNotExist(err) {
		fmt.Println("Error: Directory \"" + DirectoryPath + "\" does not exist")
		return
	}
	if !info.IsDir() {
		fmt.Println("Error: Directory \"" + DirectoryPath + "\" is not a directory")
		return
	}

	err = os.Chdir(DirectoryPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cwd, _ := os.Getwd()
	fmt.Println("Current directory:", cwd)
}

func GetPwd() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Print("Current directory:", dir, "\n")
}

func GetCP(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return fmt.Errorf("ошибка открытия исходного файла: %w", err)
	}
	defer src.Close()

	dstInfo, err := os.Stat(dstName)
	if err == nil && dstInfo.IsDir() {
		dstName = filepath.Join(dstName, filepath.Base(srcName))
	}

	dst, err := os.Create(dstName)
	if err != nil {
		return fmt.Errorf("ошибка создания файла назначения: %w", err)
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return fmt.Errorf("ошибка копирования содержимого: %w", err)
	}

	info, err := os.Stat(srcName)
	if err != nil {
		return fmt.Errorf("ошибка получения информации о файле: %w", err)
	}
	if err := os.Chmod(dstName, info.Mode()); err != nil {
		return fmt.Errorf("ошибка применения прав доступа: %w", err)
	}

	return nil
}
