package main

import "fmt"

func main() {
	// If Statment
	name := "Agus"

	if name == "Eko" {
		fmt.Println("Hello Eko")
	} else if name == "Agus" {
		fmt.Println("Ini Agus")
	} else {
		fmt.Println("Ini Bukan Eko")
	}

	// switch statment
	day := 4
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	default:
		fmt.Println("Invalid day")
	}
}
