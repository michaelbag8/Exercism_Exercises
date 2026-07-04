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
func BirdsInWeek(birdsPerDay []int, week int) int {
	// panic("Please implement the BirdsInWeek() function")
	sum := 0
	weekIdx := (week * 7) - 7
	for i := weekIdx; i < weekIdx+7; i++ {
		sum += birdsPerDay[i]
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
