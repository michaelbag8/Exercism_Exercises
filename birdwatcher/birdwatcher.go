package main

import "fmt"

//package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	sum := 0
	for _, bird := range birdsPerDay {
		sum += bird
	}
	return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	sum := 0
	switch week {
	case 1:
		for i := 0; i < 7; i++ {
			sum += birdsPerDay[i]
		}
	case 2:
		for i := 7; i < 14; i++ {
			sum += birdsPerDay[i]
		}
	case 3:
		for i := 14; i < 21; i++ {
			sum += birdsPerDay[i]
		}
	}

	return sum
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
func main() {

	birdsPerDays := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(TotalBirdCount(birdsPerDays))

	birdsPerDay := []int{2, 5, 0, 7, 4, 1, 3, 0, 2, 5, 0, 1, 3, 1}
	fmt.Println(BirdsInWeek(birdsPerDay, 2))

	birdsPerDayz := []int{2, 5, 0, 7, 4, 1}
	fmt.Println(FixBirdCountLog(birdsPerDayz))

}
