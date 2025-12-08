/*
The Cars package is used to calculate the cars produced in hours and per minute, with the total cost of making them
by using the successRate and productionRate values.
*/
package cars

var(
    //productionRate variable indicates the number of cars being assembled on the assembly line.
	productionRate int = 1000
    //The successRate variable indicates whether the number of cars being produced is working.
    successRate float64 = 90
    )
// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (float64(productionRate))*(successRate/float64(100))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return ((int(CalculateWorkingCarsPerHour(productionRate, successRate)))/int(60))
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    if carsCount <= 9{
        return uint(carsCount)*uint(10000)
    } else {
    	return ((uint(carsCount)/uint(10))*uint(95000)) + ((uint(carsCount)%uint(10))*uint(10000))
    }
}
