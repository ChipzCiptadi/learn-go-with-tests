package structs

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 5.0, Height: 10.0}, 30.0},
		{Circle{Radius: 10.0}, 62.83185307179586},
	}

	for _, test := range perimeterTests {
		got := test.shape.Perimeter()

		if got != test.want {
			t.Errorf("%#v got %f, want %f", test.shape, got, test.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 5.0, Height: 10.0}, 50.0},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Base: 12.0, Height: 6.0}, 36.0},
	}

	for _, test := range areaTests {
		got := test.shape.Area()

		if got != test.want {
			t.Errorf("%#v got %f, want %f", test.shape, got, test.want)
		}
	}
}
