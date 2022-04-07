package figuras

//ESTRUCTURA DEL CUADRADO
type Cuadrado struct {
	Lado float64
}

// METODO AREA PARA EL CUADRADO
func (cua *Cuadrado) area() float64 {
	return cua.Lado * cua.Lado
}

//METODO PERIMETRO PARA EL CUADRADO
func (cua *Cuadrado) perimetro() float64 {
	return 4 * cua.Lado
}
