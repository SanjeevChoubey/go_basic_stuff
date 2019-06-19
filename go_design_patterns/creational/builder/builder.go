package builder

type buildProcess interface {
	// Note: All these methods are exposed
	SetWheels() buildProcess
	SetSeats() buildProcess
	SetStructure() buildProcess
	GetVehicle() vehcleProduct
}

type ManufacturingDirector struct {
	builder buildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b buildProcess) {
	f.builder = b
}

// Product
type vehcleProduct struct {
	// Note: These variables are exposed to outer world
	Wheels    int
	Seats     int
	Structure string
}

// Car Builder
type CarBuilder struct {
	v vehcleProduct
}

func (c *CarBuilder) SetWheels() buildProcess {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() buildProcess {
	c.v.Seats = 4
	return c
}

func (c *CarBuilder) SetStructure() buildProcess {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) GetVehicle() vehcleProduct {
	return c.v
}

// Bike Builder
type BikeBuilder struct {
	v vehcleProduct
}

func (b *BikeBuilder) SetWheels() buildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() buildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() buildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() vehcleProduct {
	return b.v
}
