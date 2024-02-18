package main

import (
  "fmt"
)

func main() {
	num := 5
	squareOf(&num)
	fmt.Printf("Square of num %d\n",num)

	num2 := getNum()
	fmt.Println("Address returned by getNum ",num2)
	*num2 = *num2 * *num2 * *num2
	fmt.Println("Cube is ",*num2)
}

func squareOf(p *int) {
	*p = *p * *p
}

func getNum() *int {
	num := 10
	fmt.Println("Address in getNum ",&num)
	return &num
}
