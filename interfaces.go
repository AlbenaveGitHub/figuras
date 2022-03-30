package figuras

import "fmt"

type Figura interface {
	area() float64
	perimetro() float64
}

func Calculo(f Figura) {
	fmt.Println("El Calculo es: ", f)
	fmt.Println("El Area es: ", f.area())
	fmt.Println("El Perimetro es: ", f.perimetro())
}
