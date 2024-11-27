package area

import "math"

const PI = 3.1416

func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * PI
}

func Rect(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return base * altura / 2
}
