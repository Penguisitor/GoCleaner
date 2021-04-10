package clear

import (
	"fmt"
	"gocleaner/pkg/myos"
	"log"
	"os"
)

func FastClear() {

	list := [8]string{
		os.Getenv(`SystemRoot`) + `\TEMP`,
		os.Getenv("TEMP"),
		os.Getenv(`SystemRoot`) + `\Prefetch`,
		os.Getenv(`SystemRoot`) + `\SoftwareDistribution\Download`,
		`C:/ProgramData/Microsoft/Windows/WER`,
		os.Getenv(`SystemRoot`) + `\ServiceProfiles\NetworkService\AppData\Local\Microsoft\Windows\DeliveryOptimization`,
		`C:/ProgramData/Microsoft/Windows Defender/Scans/History`,
		myos.GetUserName() + `\AppData\Local\Microsoft\Windows\Explorer`,
	}

	resultSize := 0

	for _, v := range list {
		log.Println(v)
		size, err := myos.ClearFolderFiles(v)
		if err != nil {
			log.Println("[ERROR] ", err)
		}
		resultSize += size
	}
	fmt.Println("clear ", myos.ByteByMB(resultSize), " MB")
}
