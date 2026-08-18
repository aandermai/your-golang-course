package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var weight, height, bodyMassIndex float64

	fmt.Print("Введите ваш вес (кг): ")
	_, err := fmt.Scan(&weight)
	if err != nil {
		log.Fatal("Ошибка при вводе данных: ", err.Error())
	}

	fmt.Print("Введите ваш рост (см): ")
	_, err = fmt.Scan(&height)
	if err != nil {
		log.Fatal("Ошибка при вводе данных: ", err.Error())
	}

	bodyMassIndex = weight / math.Pow(height/100, 2)
	fmt.Printf("Ваш ИМТ: %.2f\n", bodyMassIndex)

	switch {
	case bodyMassIndex >= 0 && bodyMassIndex < 18.5:
		fmt.Println("Категория: Недостаточный вес")
	case bodyMassIndex >= 18 && bodyMassIndex < 25:
		fmt.Println("Категория: Нормальный вес")
	case bodyMassIndex >= 25 && bodyMassIndex < 30:
		fmt.Println("Категория: Избыточный вес")
	case bodyMassIndex >= 30:
		fmt.Println("Категория: Ожирение")
	default:
		fmt.Println("Ошибка: Указаны некорректные данные")
	}

}
