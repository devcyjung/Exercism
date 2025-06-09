package birdwatcher

func TotalBirdCount(birdsPerDay []int) int {
	s := 0
    for _, v := range(birdsPerDay) {
        s += v
    }
    return s
}

func BirdsInWeek(birdsPerDay []int, week int) int {
    s := 0
	for _, v := range(birdsPerDay[7 * (week-1):7 * week]) {
        s += v
    }
    return s
}

func FixBirdCountLog(birdsPerDay []int) []int {
	for i := range(birdsPerDay) {
        if i % 2 == 1 {
            continue
        }
        birdsPerDay[i]++
    }
    return birdsPerDay
}
