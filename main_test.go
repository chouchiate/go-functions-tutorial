// cmd: go test -v -cover=true main_test.go main.go
package main

import (
	"testing"

)

func TestSimpleFunction(t *testing.T) {
	expected := "Chia Chou"
	result := SimpleFunction("Chia","Chou")

	if expected != result {
		t.Error("The result does not match what we expected")
	}
}

func TestSimpleFuncWithError(t *testing.T){
	expected := "Chia Chou"
	result, err := SimpleFuncWithError("Chia","Chou")
	if err != nil {
		t.Error("Our Function Threw an Error")
	}
	if expected != result {
		t.Error("The result does not match what we expected")
	}
}