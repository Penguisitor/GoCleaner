package myos

import (
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
)

// GetUserName - получить имя текущего авторизированного пользователя
func GetUserName() string {
	cur, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return cur.HomeDir
}

// ClearFolderFiles - очистить папку, получив в замен сумму размера всех обработанных файлов
func ClearFolderFiles(pathToFolder string) (int, error) {
	sizeData := 0
	dir, err := ioutil.ReadDir(pathToFolder)
	if err != nil {
		return sizeData, err
	}
	for _, d := range dir {
		sizeData += int(d.Size())
		os.RemoveAll(path.Join([]string{pathToFolder, d.Name()}...))
	}
	return sizeData, nil
}

// ByteByMB - перводит байты в мегабайты
func ByteByMB(Byte int) int {
	return (Byte / 1024) / 1024
}
