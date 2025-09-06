package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulanagn ke ", counter)
		counter++
	}

	fmt.Println("Selesai")
}
