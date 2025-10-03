# 📂 Go File Manager CLI

A simple **CLI file manager** written in Go.  
This project is my first step into Go development.  
Currently supports basic file operations.

---

## 🚀 Commands

- `cd <path>` — change directory
- `pwd` — show current working directory
- `cp <file> <destination>` — copy a file
- `mv <file> <destination>` — move a file
- `rm <file>` — remove (delete) a file
- `mkdir <dirname>` — create a new directory
- `find <root> <filename>` — search for a file starting from the current directory

---

## ▶️ Usage

```bash
git clone https://github.com/AKUMA365/TinyFileManager.git
cd TinyFileManager
go build -o filemgr
./filemgr

