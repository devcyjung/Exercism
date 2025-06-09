package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	s := 0
    for _, v := range(birdsPerDay) {
        s += v
    }
    return s
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    s := 0
	for _, v := range(birdsPerDay[7 * (week-1):7 * week]) {
        s += v
    }
    return s
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range(birdsPerDay) {
        if i % 2 == 1 {
            continue
        }
        birdsPerDay[i]++
    }
    return birdsPerDay
}
