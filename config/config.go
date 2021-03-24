package config

// ConfigureClear структура, содержащая информацию о функциях отчистки приложения
type ConfigureClearStruct struct {
	// Временные файлы работы Windows
	WindowTemp bool
	// Временные файлы работы программ пользователя
	UserTemp bool
	// Временные файлы модуля windows Prefetch, ускоряющий загрузку системы при включении
	Prefetch bool
	// При обновлении Windows скачивает в отдельную папку все необходимые файлы
	// для установки и не удаляет за собой.
	WindowUpdate bool
	// Эскизы Windows, это копии файлов изображений, видео и документов
	// для их быстрого отображения в окне проводника
	ExplorerTemp bool
	// Файлы оптимизации доставки - это файлы полученные из центра обновления Windows или MicrosoftStore,
	// которые храняться у вас для более быстрой передачи их между компьютера в локальной сети
	DeliveryOptimizationTemp bool
	// Временные файлы, логи WindowsDefender
	WindowsDefenderTemp bool
	// Временные файлы, логи отчетов об ошибках Windows
	WerTemp bool
}

// ConfigureClear структура, содержащая информацию о функциях отчистки приложения
var ConfigureClear ConfigureClearStruct = ConfigureClearStruct{
	true, true, true, true, true, true, true, true,
}

const WerDirTemp = "C:/ProgramData/Microsoft/Windows/WER"

const WindowsDefenderDirTemp = "C:/ProgramData/Microsoft/Windows Defender/Scans/History"

const DeliveryOptimizationDirTemp = "C:/Windows/ServiceProfiles/NetworkService/AppData/Local/Microsoft/Windows/DeliveryOptimization"

// Путь до временных файлов windows
const WindowsDirTemp = "C:/Windows/Temp"

// Путь до временных файлов пользователя без учета имени пользователя
const UserDirTemp = "/AppData/Local/Temp"

// Путь до временных файлов Prefetch
const PrefetchDirTemp = "C:/Windows/Prefetch"

// Путь до временных файлов обновления windows
const WindowDirUpdate = "C:/Windows/SoftwareDistribution/Download"

// Путь до временных файлов экскизов
const ExplorerDirTemp = "/AppData/Local/Microsoft/Windows/Explorer"
