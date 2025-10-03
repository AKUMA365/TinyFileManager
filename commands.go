package main

import (
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
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

	case "mv":
		if len(parts) < 3 {
			fmt.Println("Error: Command \"mv\" requires at least 3 arguments")
			return
		}
		GetMV(parts[1], parts[2])

	case "rm":
		if len(parts) < 2 {
			fmt.Println("Error: Command \"rm\" requires at least 2 arguments")
		}
		GetRM(parts[1])
	case "mkdir":
		if len(parts) < 2 {
			fmt.Println("Error: Command \"mkdir\" requires at least 2 arguments")
		}
		GetMKdir(parts[1])
	case "find":
		if len(parts) < 2 {
			fmt.Println("Error: Command \"find\" requires at least 2 arguments")
		}
		GetFind(parts[1], parts[2])
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

func GetMV(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	err = os.Remove(src)
	if err != nil {
		return err
	}

	return nil
}

func GetRM(FileName string) error {

	err := os.Remove(FileName)
	if err != nil {
		return err
	}

	return nil
}

func GetMKdir(DirName string) error {
	err := os.Mkdir(DirName, 0777)
	if err != nil {
		return err
	}

	return nil
}
func roots() []string {
	if runtime.GOOS == "windows" {
		var r []string
		for c := 'A'; c <= 'Z'; c++ {
			drive := fmt.Sprintf("%c:\\", c)
			if _, err := os.Stat(drive); err == nil {
				r = append(r, drive)
			}
		}
		return r
	}
	// unix-like
	return []string{"/"}
}

func GetFind(root, target string) (string, error) {
	var found string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() == target {
			found = path
			fmt.Printf("✅ File found: %s\n", path) // Вывод прямо в момент нахождения
			return filepath.SkipDir                // можно остановить после первого найденного
		}
		return nil
	})

	if err != nil {
		return "", err
	}

	if found == "" {
		fmt.Println("❌ File not found.")
		return "", nil
	}

	return found, nil
}
