package main

import (
	"fmt"
	"gocleaner/internal/clear"
)

func main() {

	for true {
		var command string
		fmt.Println("")
		fmt.Println("Добро пожаловать в программу gocleaner!")
		fmt.Println("")
		fmt.Println("1. Fast clear - Очистить систему от остаточных и временных файлов")
		fmt.Println("2. Exit - Выход")
		fmt.Scanf("%s\n", &command)
		if command == "1" {
			clear.StartClear()
		}
		if command == "2" {
			return
		}
	}
}
