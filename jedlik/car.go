package jedlik
// 1. Added the 'distance' and 'battery' fields to track the car's state.
type Car struct {
	speed        int
	batteryDrain int
	battery      int
	distance     int // Tracks total meters driven
}

// NewCar creates a new car with a full 100% battery.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:        speed,
		batteryDrain: batteryDrain,
		battery:      100,
		distance:     0,
	}
}

// 2. Method with a pointer receiver (*Car) to update the fields directly.

