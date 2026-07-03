package jedlik

import "fmt"

// Drive updates distance and battery
func (c *Car) Drive() {
    if c.battery < c.batteryDrain {
        return
    }

    c.distance += c.speed
    c.battery -= c.batteryDrain
}

// DisplayDistance returns formatted distance string
func (c *Car) DisplayDistance() string {
    return fmt.Sprintf("Driven %d meters", c.distance)
}

// DisplayBattery returns formatted battery string
func (c *Car) DisplayBattery() string {
    return fmt.Sprintf("Battery at %d%%", c.battery)
}

// CanFinish checks if car can complete track
func (c *Car) CanFinish(trackDistance int) bool {
    maxDrives := c.battery / c.batteryDrain
    maxDistance := maxDrives * c.speed

    return maxDistance >= trackDistance
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
