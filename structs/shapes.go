package structs

import "math"

// Shape an interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle just a basic geometry
type Rectangle struct {
	Width  float64
	Height float64
}

// Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle just a round geometry
type Circle struct {
	Radius float64
}

// Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter of Circle
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

// Triangle just a pointy geomtry
type Triangle struct {
	Base   float64
	Height float64
}

// Area of Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}

// Perimeter of Triangle (unimplemented)
func (t Triangle) Perimeter() float64 {
	return 0
}

// Perimeter some geometry code to calculate the perimeter of a rectangle given a width and height
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
