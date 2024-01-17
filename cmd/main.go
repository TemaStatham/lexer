package main

import (
	"fmt"
	"lexer/pkg/lexer"
	"os"
)
 
func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Неверное количество аргументов командной строки: ", len(args))
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println("Ошибка открытия файла: ", err)
		return
	}
	defer file.Close()

	lexer.Print( lexer.Parse(file))
	
}
