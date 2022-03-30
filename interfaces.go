package figuras

import "fmt"

type figura interface {
	area() float64
	perimetro() float64
}

func Calculo(f figura) {
	fmt.Println("El Calculo es: ", f)
	fmt.Println("El Area es: ", f.area())
	fmt.Println("El Perimetro es: ", f.perimetro())
}
