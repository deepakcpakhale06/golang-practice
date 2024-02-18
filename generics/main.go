package main

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

type superFloat float64

type allowedTypes interface {
	constraints.Float | constraints.Signed
}

func main(){
	fmt.Println(minInt(5,6))
	fmt.Println("========================")
	fmt.Println(minFloat(5.989,5.990))
	fmt.Println("========================")
	fmt.Println(minGeneric[int](5,6))
	fmt.Println("========================")
	fmt.Println(minGeneric[float64](5.989,5.990))
	fmt.Println("========================")
	var v1, v2 superFloat = 0.9, 0.10
	fmt.Println(minGeneric[superFloat](v1,v2))
	fmt.Println("========================")
}

func minGeneric[T allowedTypes ](a T, b T) T {
	fmt.Println("Comparing using Generic Function")
	if a < b {
		return a
	}

	return b
}

func minInt(a int, b int) int {

	fmt.Println("Comparing Integers")
	if a < b {
		return a
	}

	return b

}

func minFloat(a float64, b float64) float64 {

	fmt.Println("Comparing Floats")
	if a < b {
		return a
	}

	return b

}