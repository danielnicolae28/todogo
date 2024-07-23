package main

import (
	"fmt"
)

func main() {
	var inputToDo string
	todos := []string{}
	// todos := make([]string, 100)
	fmt.Println("To Do List")
	for {
		fmt.Print("Write To Do: ")
		fmt.Scan(&inputToDo)
		length := len(inputToDo)
		if length < 3 {
			fmt.Println("Input is to short")
		}
		todos = append(todos, inputToDo)

		fmt.Println(todos)
	}

}
