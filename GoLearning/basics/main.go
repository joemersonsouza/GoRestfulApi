package main

import (
	"fmt"
	"golearning/basics/middlepkg"
)

func main() {
	var projectName string
	fmt.Println("What is the project name")
	fmt.Scan(&projectName)

	fmt.Printf("Starting the project %v\n", projectName)

	middlepkg.Point = "Cheguei\n"
	middlepkg.PrintMiddle()

	var bookings []string

	bookings = append(bookings, "Oi Jose")

	for _, booking := range bookings {
		fmt.Println(booking)
	}

	var myMap = make(map[int]string)
	myMap[1] = "Força"
	println(myMap[1])
	end()
}
