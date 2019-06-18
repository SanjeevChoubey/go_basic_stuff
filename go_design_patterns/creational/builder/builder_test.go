package builder

import (
	"testing"
)

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}
	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()
	car := carBuilder.GetVehicle()

	if car.Structure != "Car" {
		t.Errorf("Structure of Car was supposed to be Car but it is %s ", car.Structure)
	}

	if car.Seats != 4 {
		t.Errorf("Seat in Car was supposed to be 4 but it is %d", car.Seats)
	}

	if car.Wheels != 4 {
		t.Errorf("Wheels of Car was supposed to be 4 but it is %d ", car.Wheels)
	}

}
