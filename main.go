package main

import (
	"errors"
	"fmt"
)

func SimpleFunction(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func SimpleFuncWithError(firstName string, lastName string) (string,error) {
	if len(firstName) <= 1 {
		return "", errors.New("First Name Must Be At Least 2 letters long")
	}

	return firstName + " " + lastName, nil
}

func addOne() func() int {
	var x int
	return func() int {
		x++
		return x + 1 
	}
}

// Method

type Employee struct {
	Name string

}

func (e *Employee) UpdateName(newName string) {
	e.Name = newName
}

func (e *Employee) PrintName () {
	fmt.Println(e.Name)
}

func main() {
	fmt.Println("Go functions Tutorial")

	fullName := SimpleFunction("Chia","Chou")
	fmt.Println(fullName)

	fullName2, err := SimpleFuncWithError("","Chou")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fullName2)

	myFunc := addOne()
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())
	fmt.Println(myFunc())

	var employee Employee
	employee.Name = "Chia"
	employee.UpdateName("Chou")
	employee.PrintName()
}