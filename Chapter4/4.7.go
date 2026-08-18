/*
	Практика -- поиск по названию товара
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	keyboardQualities   = "Клавиатура JZ9: 19200"
	headphonesQualities = "Наушники N45: 9600"
	phoneQualities      = "Смартфон S10: 55000"
)

func main() {
	fmt.Print("Введите название товара: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		log.Fatal("Ошибка ввода данных", err.Error())
	}

	productName := scanner.Text()
	keyboardName := strings.Split(keyboardQualities, ":")[0]
	headphonesName := strings.Split(headphonesQualities, ":")[0]
	phoneName := strings.Split(phoneQualities, ":")[0]

	switch {
	case strings.Contains(strings.ToLower(keyboardName), strings.ToLower(productName)):
		fmt.Println(keyboardQualities)
	case strings.Contains(strings.ToLower(headphonesName), strings.ToLower(productName)):
		fmt.Println(headphonesQualities)
	case strings.Contains(strings.ToLower(phoneName), strings.ToLower(productName)):
		fmt.Println(phoneQualities)
	default:
		fmt.Printf("Товар \"%s\" не найден.\n", productName)
	}
}
