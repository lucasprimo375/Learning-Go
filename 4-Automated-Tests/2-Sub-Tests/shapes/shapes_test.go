package shapes

import (
	"math"
	"testing"
)

func TestArea(t *testing.T) {
	t.Run("Rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 12}

		expectedArea := float64(120)

		receivedArea := rectangle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Received area %f is different than the expected %f", receivedArea, expectedArea)
		}
	})

	t.Run("Circle", func(t *testing.T) {
		circle := Circle{10}

		expectedArea := float64(math.Pi * 100)

		receivedArea := circle.Area()

		if expectedArea != receivedArea {
			t.Fatalf("Received area %f is different than the expected %f", receivedArea, expectedArea)
		}
	})
}
