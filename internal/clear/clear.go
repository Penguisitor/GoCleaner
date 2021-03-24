package clear

import (
	"fmt"
	"gocleaner/config"
	"gocleaner/pkg/myos"
	"log"
	"os/exec"
)

func StartClear() {

	resultSize := 0

	log.Println("Window temp is cleared...")
	if config.ConfigureClear.WindowTemp {
		size, err := myos.ClearFolderFiles(getWindowTemp())
		if err != nil {
			log.Println("[ERROR] Window temp: ", err)
		}
		resultSize += size
	}

	log.Println("Window update temp is cleared...")
	if config.ConfigureClear.WindowUpdate {
		size, err := myos.ClearFolderFiles(gewWindowUpdate())
		if err != nil {
			log.Println("[ERROR] Window update temp: ", err)
		}
		resultSize += size
	}

	log.Println("User temp is cleared...")
	if config.ConfigureClear.UserTemp {
		size, err := myos.ClearFolderFiles(getUserTemp())
		if err != nil {
			log.Println("[ERROR] User temp: ", err)
		}
		resultSize += size
	}

	log.Println("Prefetch temp is cleared...")
	if config.ConfigureClear.Prefetch {
		size, err := myos.ClearFolderFiles(getPrefetchTemp())
		if err != nil {
			log.Println("[ERROR] Prefetch temp: ", err)
		}
		resultSize += size
	}

	log.Println("Explorer temp is cleared...")
	err := exec.Command("cmd", "/C", "taskkill /IM explorer.exe /F").Run()
	if err != nil {
		log.Println("[ERROR] Explorer temp: ", err)
	}
	if config.ConfigureClear.ExplorerTemp {
		size, err := myos.ClearFolderFiles(gerExplorerDirTemp())
		if err != nil {
			log.Println("[ERROR] Explorer temp: ", err)
		}
		resultSize += size
	}
	err = exec.Command("cmd", "/C", "start explorer.exe").Run()
	if err != nil {
		log.Println("[ERROR] Explorer temp: ", err)
	}

	log.Println("DeliveryOptimization temp is cleared...")
	if config.ConfigureClear.ExplorerTemp {
		size, err := myos.ClearFolderFiles(gerDeliveryOptimizationDirTemp())
		if err != nil {
			log.Println("[ERROR] DeliveryOptimization temp: ", err)
		}
		resultSize += size
	}

	log.Println("Windows Defender temp is cleared...")
	if config.ConfigureClear.ExplorerTemp {
		size, err := myos.ClearFolderFiles(getWindowsDefenderDirTemp())
		if err != nil {
			log.Println("[ERROR] Windows Defender temp: ", err)
		}
		resultSize += size
	}

	log.Println("Wer temp is cleared...")
	if config.ConfigureClear.ExplorerTemp {
		size, err := myos.ClearFolderFiles(getWerDirTemp())
		if err != nil {
			log.Println("[ERROR] Wer temp: ", err)
		}
		resultSize += size
	}

	fmt.Println("clear ", byteByMB(resultSize), " MB")
	fmt.Println("TEMP is cleared")
}
