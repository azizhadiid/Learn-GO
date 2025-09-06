package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var aziz Customer
	aziz.Name = "Aziz Alhadiid"
	aziz.Address = "Jambi"
	aziz.Age = 20
	fmt.Println(aziz)
	fmt.Println(aziz.Name)

	joko := Customer{
		Name:    "Joko",
		Address: "Jakarta",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"Budi", "Depok", 39}
	fmt.Println(budi)
}
