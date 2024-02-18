package main

import (
	"math"
)

type shape interface {
	area() float64
	circumference() float64
}

type shapeV2 interface {
	areaV2() string
	circumferenceV2() float64
}

type square struct {
	side float64
}

type circle struct {
	radius float64
}

//square methods
func (s square) area() float64 {
   return s.side * s.side
}

func (s square) circumference() float64 {
   return s.side * 4
}

func (s square) areaV2() string {
	return "areaV2() Method is called for Square shape"
 }

func (s square) circumferenceV2() float64 {
	return s.side + s.side + s.side + s.side
 }

//circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) areaV2() string {
	return "areaV2() Method is called for Circle shape"
}

func (c circle) circumferenceV2() float64 {
	return 2 * math.Pi * c.radius
}



