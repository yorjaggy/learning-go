package main

import "fmt"

func main() {
	//Testing pointers and references
	varWithValue := 3312

	varWithReference := &varWithValue

	fmt.Printf("Memory Address %v - Value in var %v\n", varWithReference, *varWithReference)

	fmt.Println("Changing the value using the pointer")
	*varWithReference = 3313
	fmt.Printf("Memory Address %v - Value in var %v\n", varWithReference, *varWithReference)

	fmt.Println("Changing the value using the var name")
	varWithValue = 3314
	fmt.Printf("Memory Address %v - Value in var %v\n", varWithReference, *varWithReference)

	//&varWithValue = 3315 //This is not permitted by go
	//&varWithValue := 3315 //Neither this

	otherVar := 3315

	/*
		varWithReference = &otherVar

		This is not accepted, due to otherVar is a reference and
		the pointer in varWithReference is int type
	*/
	fmt.Println("Changing the value using other reference")

	varWithReference = &otherVar
	fmt.Printf("Memory Address %v - Value in var %v\n", varWithReference, *varWithReference)

	fmt.Printf("Initial var is -> %v\n", varWithValue)
}
