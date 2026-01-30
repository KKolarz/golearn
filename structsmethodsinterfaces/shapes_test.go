package shapes

import "testing"

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{10.0, 5.0}

	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}

}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()

		if got != hasArea {
			t.Errorf("%#v Got %g, want %g", shape, got, hasArea)
		}
	}

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{10, 6}, hasArea: 30},
	}

	for _, tt := range areaTests {
		checkArea(t, tt.shape, tt.hasArea)
	}

	// t.Run("Rectangles:", func(t *testing.T) {
	// 	rectangle := Rectangle{10.0, 5.0}

	// 	checkArea(t, rectangle, 50.0)
	// })

	// t.Run("Circles:", func(t *testing.T) {
	// 	circle := Circle{10.0}

	// 	checkArea(t, circle, 314.1592653589793)
	// })
}
