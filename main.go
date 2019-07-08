package goarea

import "math"

//Pi é PI
const Pi = 3.1415

//Circ calcula a área de um círculo
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

//Rect calcula a área de um retângulo
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return base * altura / 2
}
