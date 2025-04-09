package main

import (
	"fmt"
	"lab_1/generate"
)

func main() {
	// Генерация всех данных
	if err := generate.GenerateAllData(); err != nil {
		fmt.Printf("Ошибка генерации данных: %v\n", err)
		return
	}
	fmt.Println("Данные успешно сгенерированы в папке data/")
}