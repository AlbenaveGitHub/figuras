package figuras

import "fmt"

//INTERFASE
type Figura interface {
	area() float64
	perimetro() float64
}

// FUNCION PARA MOSTRAR LOS RESULTADOS
func Calculo(f Figura) {
	fmt.Println("Número a calcular:", f)
	fmt.Println("Area calculada:", f.area())
	fmt.Println("Perimetro calculado:", f.perimetro())
}
