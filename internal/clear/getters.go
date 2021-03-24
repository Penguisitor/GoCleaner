package clear

import (
	"gocleaner/config"
	"gocleaner/pkg/myos"
)

// getWindowsDefenderDirTemp - получить путь до временных файлов отчетов windows
func getWerDirTemp() string {
	return config.WerDirTemp
}

// getWindowsDefenderDirTemp - получить путь до временных файлов windows defender
func getWindowsDefenderDirTemp() string {
	return config.WindowsDefenderDirTemp
}

// DeliveryOptimizationDirTemp - получить путь до временных файлов оптимизации загрузки
func gerDeliveryOptimizationDirTemp() string {
	return config.DeliveryOptimizationDirTemp
}

// GerExplorerDirTemp - получить путь до временных файлов экскизов (мультимедия)
func gerExplorerDirTemp() string {
	return myos.GetUserName() + config.ExplorerDirTemp
}

// GewWindowUpdate - получить путь до остаточных файлов обновления windows
func gewWindowUpdate() string {
	return config.WindowDirUpdate
}

// ByteByMB - перводит байты в мегабайты
func byteByMB(Byte int) int {
	return (Byte / 1024) / 1024
}

// GetWindowTemp - получить путь до временных файлов Prefetch (это компонент windows, который ускоряет
// загрузку систему)
func getPrefetchTemp() string {
	return config.PrefetchDirTemp
}

// GetWindowTemp - получить путь до временных файлов windows
func getWindowTemp() string {
	return config.WindowsDirTemp
}

// GetUserTemp - получить путь до временных файлов пользователя
func getUserTemp() string {
	return myos.GetUserName() + config.UserDirTemp
}
