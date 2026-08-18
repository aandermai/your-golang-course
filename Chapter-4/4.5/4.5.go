package main

import (
	"fmt"
	"log"
)

func main() {
	var dayTime int

	fmt.Print("Введите время в часах: ")
	_, err := fmt.Scan(&dayTime)
	if err != nil {
		log.Fatal("Ошибка считывания данных с терминала: ", err.Error())
	}

	switch {
	case dayTime >= 6 && dayTime < 12:
		fmt.Printf("Сейчас %dч. - утро\n", dayTime)
	case dayTime >= 12 && dayTime < 18:
		fmt.Printf("Сейчас %dч. - день\n", dayTime)
	case dayTime >= 18 && dayTime < 23:
		fmt.Printf("Сейчас %dч. - вечер\n", dayTime)
	case dayTime == 23 || dayTime >= 0:
		fmt.Printf("Сейчас %dч. - ночь\n", dayTime)
	default:
		fmt.Println("Неверно задано время")
	}
}
