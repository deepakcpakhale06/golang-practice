package main

import(
	"fmt"
)

type allowedShapes interface {
	// Can not use shape | shapeV2, its allowed only for basic datatypes such as int | int 32 | float64 etc.
	shape
}

func main(){
	s1 := square{
		side : 10.0,
	}
	getShapeDetails(s1)
}

func getShapeDetails [T allowedShapes] (s T) {
	fmt.Println("Area is ",s.area())
	fmt.Println("Circumference is ",s.circumference())
	//fmt.Println("Area Calculated via V2 is ",s.areaV2())
	//fmt.Println("Circumference calculated via V2 is ",s.circumferenceV2())
}