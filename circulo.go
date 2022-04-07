package figuras

import "math"

//ESTRUCTURA DEL CIRCULO
type Circulo struct {
	Radio float64
}

// METODO AREA PARA EL CIRCULO
func (cir *Circulo) area() float64 {
	return math.Pi * (cir.Radio * cir.Radio)
}

//METODO PERIMETRO PARA EL CIRCULO
func (cir *Circulo) perimetro() float64 {
	return math.Pi * cir.Radio
}
