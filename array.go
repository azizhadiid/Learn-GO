package main

import (
	"fmt"
)

func main() {
	var names [3]string
	names[0] = "Aziz"
	names[1] = "Al"
	names[2] = "Hadiid"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		95,
	}

	fmt.Println(" ")
	fmt.Println(values)

	slicess := [...]string{"Aziz", "Al", "Hadiid", "Eko", "Joko", "Agus", "Hasby"}

	slice1 := slicess[4:6]
	fmt.Println(slice1)

	slice2 := slicess[4:]
	fmt.Println(slice2)

	slice3 := slicess[:4]
	fmt.Println(slice3)

	slice4 := slicess[:]
	fmt.Println(slice4)

	days := [...]string{"Aziz", "Al", "Hadiid", "Eko", "Joko", "Agus", "Hasby"}
}
