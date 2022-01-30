package structsmethodsinterfaces

import (
	"reflect"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func getType(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return "*" + t.Elem().Name()
	} else {
		return t.Name()
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape             Shape
		expectedPerimeter float64
	}{
		{shape: Rectangle{Width: 10.0, Height: 10.0}, expectedPerimeter: 40.0},
		{shape: Circle{Radius: 10.0}, expectedPerimeter: 62.83185307179586},
		{shape: Triangle{Base: 10.0, Height: 10.0, SideA: 10.0, SideC: 10.0}, expectedPerimeter: 30},
	}

	for _, tt := range perimeterTests {
		Convey("Should return perimeter of "+getType(tt.shape), t, func() {
			got := tt.shape.Perimeter()
			So(got, ShouldEqual, tt.expectedPerimeter)
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape        Shape
		expectedArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, expectedArea: 72.0},
		{shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}

	for _, tt := range areaTests {
		Convey("Should return area of "+getType(tt.shape), t, func() {
			got := tt.shape.Area()
			So(got, ShouldEqual, tt.expectedArea)
		})
	}
}
