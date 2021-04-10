package clear

import (
	"fmt"
	"gocleaner/pkg/myos"
	"log"
	"os"
	"os/exec"
)

func ClearDownloads() {
	path := os.Getenv("USERPROFILE") + `\Downloads`

	resultSize := 0

	log.Println(path)

	size, err := myos.ClearFolderFiles(path)
	if err != nil {
		log.Println("[ERROR] ", err)
	}
	resultSize += size
	fmt.Println("clear ", myos.ByteByMB(resultSize), " MB")
}

func ClearBasket() {
	log.Println("Clear RecycleBin")
	cmd := exec.Command(`powershell.exe -command Clear-RecycleBin`)
	_, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
}
