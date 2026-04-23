package main

import "fmt"

func main() {

	//Hola mundo

	fmt.Println("Hola, Go!")

	/*
		Comentario
		de
		varias
		lineas
	*/

	//variables

	var myString string = "Esto es una cadena de texto"
	fmt.Println(myString)

	myString = "Aqui cambio el valor de la cadena de texto"
	fmt.Println(myString)

	// myString = 6 Error

	var myString2 = "Esto es otra cadena de texto"
	fmt.Println(myString2)

	var myInt int = 7
	myInt = myInt + 4
	fmt.Println(myInt)

	fmt.Printf("%s %d", myString, myInt)
	fmt.Println(myString, myInt)

}
