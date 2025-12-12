package birdwatcher

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
    // sum is intiliazed with the value zero.
    sum := 0
	for i:= 0; i < len(birdsPerDay); i++{
        sum += birdsPerDay[i]
    }
    return sum
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
    // totalBirds is intiliazed with the value zero.
    totalBirds := 0
    // startIndex is the variable that decides the starting index for the loop.
	startIndex := (week - 1) * 7
    // endIndex is the variable that decides the ending index for the loop.
    endIndex := startIndex + 7
    for i := startIndex; i < endIndex; i++{
        if i>= len(birdsPerDay){
            break
        }
        totalBirds += birdsPerDay[i]
    }
    return totalBirds
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i:= 0; i < len(birdsPerDay); i += 2{
        birdsPerDay[i] = (birdsPerDay[i] + 1)
    }
    return birdsPerDay
}
