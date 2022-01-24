// Package birdwatcher keeps track of how many birds have visited your garden
package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	var birds int

	for _, count := range birdsPerDay {
		birds += count
	}

	return birds
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	var birds int

	for ctr := (7 * week) - 7; ctr < len(birdsPerDay); ctr++ {
		if ctr < (7 * week) {
			birds += birdsPerDay[ctr]
		}
	}
	return birds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i, count := range birdsPerDay {
		if i%2 == 0 {
			birdsPerDay[i] = count + 1
		} else {
			birdsPerDay[i] = count
		}
	}

	return birdsPerDay
}
