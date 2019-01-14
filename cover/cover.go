package cover

import "math"

func Circle(r float64) float64 {
	return r * r * math.Pi
}

func Triangle(base, height float64) float64 {
	return base * height / 2
}

func Rect(width, height float64) float64 {
	return width * height
}

func Square(width float64) float64 {
	return width * width
}
