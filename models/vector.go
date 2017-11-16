package models

type Vector2 = [2]float64

func NewVector2(x, y float64) Vector2 {
	return Vector2{x, y}
}
