package main

import "fmt"

func main() {
	age := 28 // regular variable

	agePointer := &age

	fmt.Println("Age: ", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
