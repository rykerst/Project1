package main

import "fmt"

func main() {
	var text int
	_, err := fmt.Scan(&text)
	if err != nil {
		fmt.Printf("Ошибка: %s", err)
		return
	}
	fmt.Printf("Введенное число: %d", text)
}
