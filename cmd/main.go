package main

import (
	"fmt"
	"lexer/pkg/lexer"
	"os"
)
 
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Printf("Неверное количество аргументов командной строки: %d", len(args))
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	lexer.Print( lexer.Parse(file))
	
}
