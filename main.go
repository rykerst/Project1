package main

import "fmt"

func main() {
	var text int
	_, err := fmt.Scan(&text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(text)
}
