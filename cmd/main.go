package main

import (
	"fmt"
	"gocleaner/internal/clear"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

type SysInfo struct {
	Hostname string `bson:hostname`
	Platform string `bson:platform`
	CPU      string `bson:cpu`
	Cores    int    `bson:cores`
	TotalRAM uint64 `bson:totalram`
	UsedRAM  uint64 `bson:usedram`
	PRAM     uint64 `bson:pram`
	Disk     uint64 `bson:disk`
}

func info() {
	hostStat, _ := host.Info()
	cpuStat, _ := cpu.Info()
	cpuCores, _ := cpu.Counts(true)
	vmStat, _ := mem.VirtualMemory()
	diskStat, _ := disk.Partitions(true)

	// test, _ := disk.Partitions(true)
	// fmt.Println(test)

	info := new(SysInfo)

	info.Hostname = hostStat.Hostname
	info.Platform = hostStat.Platform
	info.CPU = cpuStat[0].ModelName
	info.Cores = cpuCores
	info.TotalRAM = vmStat.Total / 1024 / 1024
	info.UsedRAM = vmStat.Used / 1024 / 1024
	info.PRAM = (info.UsedRAM * 100) / info.TotalRAM

	fmt.Println("HostName: ", info.Hostname)
	fmt.Println("Platform: ", info.Platform)
	fmt.Println("CPU: ", info.CPU, info.Cores, "Cores")
	fmt.Println("RAM: ", info.UsedRAM, "/", info.TotalRAM, info.PRAM, "%")
	for _, value := range diskStat {
		v, _ := disk.Usage(value.Device)
		fmt.Println(v.Path, "", v.Used/1024/1024, "/", v.Total/1024/1024, (v.Used*100)/v.Total, "%")
	}

}

func main() {

	for true {
		var command string
		fmt.Println("")
		fmt.Println("Добро пожаловать в программу GoCleaner!")
		fmt.Println("")
		info()
		fmt.Println("")
		fmt.Println("1. Clear temp - Очистить систему от остаточных и временных файлов")
		fmt.Println("2. Clear downloads - Очистить загрузки")
		fmt.Println("3. Clear basket - Очистить корзину")
		fmt.Println("4. Exit - Выход")
		fmt.Scanf("%s\n", &command)
		if command == "1" {
			clear.FastClear()
		}
		if command == "2" {
			clear.ClearDownloads()
		}
		if command == "3" {
			clear.ClearBasket()
		}
		if command == "4" {
			return
		}
	}
}
