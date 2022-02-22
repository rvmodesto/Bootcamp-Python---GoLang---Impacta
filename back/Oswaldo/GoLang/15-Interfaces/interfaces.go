package main

import (
	"fmt"
	"math"
)

//sytax interface
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f", f.area())
}

type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64{
	return r.altura * r.largura
}


type circulo struct {
	raio float64
}

func (c circulo) area() float64 {
	return math.Pi * c.raio * c.raio
	// return math.Pi * math.Pow(c.raio)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{10}
	escreverArea(c)
}