package main

import "fmt"

func main() {

	/* declaring with var (explicit type)  */
	var name string = "GoLang"
	var age int = 25
	var pi float64 = 3.1415
	var isReady bool = true

	fmt.Printf("%s %d %.3f %t \n", name, age, pi, isReady)

	/* declaring with var (Type inference) */
	var monument = "Taj Mahal"

	fmt.Println(monument)

	var a, b, c int = 1, 2, 3
	x, y, z := "Go", 3.14, true

	fmt.Println(a, b, c, x, y, z)

	var (
		username string = "admin"
		password string = "1234"
		port     int    = 8080
	)

	fmt.Println(username, password, port)

	const firstname = "Ritika"
	const lastname = "Agrawal"
	const fullname = firstname + " " + lastname

	fmt.Println(fullname)
}
