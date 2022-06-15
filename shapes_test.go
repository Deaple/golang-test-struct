package main

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := rectangle.Perimeter()

	want := 40.0

	if got != want{
		t.Errorf("got %2.f, wanted %2.f",got,want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10},314.1592653589793},
		{Triangle{Base: 12,Height: 6}, 36.0},
	}

	for _, tt := range areaTests{
		got := tt.shape.Area()
		if got != tt.want{
			t.Errorf("got %g, wanted %g", got, tt.want)
		}
	}

	checkArea := func (t testing.TB, shape Shape, want float64)  {
		t.Helper()
		got := shape.Area()
		if got != want{
			t.Errorf("got %g, wanted %g", got, want)
		}
	}

	t.Run("rectangles", func (t *testing.T)  {
		rectangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rectangle, want)	
	})

	t.Run("circles", func(t *testing.T){
		circle := Circle{10}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
	
}