package main

import (
	"fmt"
	"reflect"
)

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

	fmt.Println(reflect.TypeOf(myInt))

	var myFloat = 6.5
	fmt.Println(myFloat)
	fmt.Println(reflect.TypeOf(myFloat))

	fmt.Println(float64(myInt) + myFloat)

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	//Variable declara e inicializada de manera abreviada

	myString3 := "Esto es una cadena de texto abreviada e inicializada"
	fmt.Println(myString3)

	//Constante

	const myConst = "Esto es una constante"
	fmt.Println(myConst)

	// Control de flujo

	myInt = 11
	myString = ""

	if myInt == 10 && myString == "Hola" {
		fmt.Println("El valor es 10")
	} else if myInt == 11 || myString == "Hola" {
		fmt.Println("El valor es 11")
	} else {
		fmt.Println("El valor no es 10")
	}

	// Array

	var myArray [3]int
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	// myArray[3] = 3 Error

	fmt.Println(myArray)
	fmt.Println(myArray[2])
	// fmt.Println(myArray[3]) Error

	// Map

	myMap := make(map[string]int)
	myMap["Eduardo"] = 36
	myMap["Laura"] = 35
	myMap["Juan"] = 24
	fmt.Println(myMap["Eduardo"])

}
