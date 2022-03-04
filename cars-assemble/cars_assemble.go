package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	var workingCars float64 = (successRate / 100) * float64(productionRate)

	return workingCars
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	var workingCars float64 = CalculateWorkingCarsPerHour(productionRate, successRate)
	var workingCarsPerMinute = workingCars / 60
	return int(workingCarsPerMinute)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	var groupOfTen int = carsCount / 10
	var units int = carsCount - (groupOfTen * 10)
	var totalCost uint = uint((groupOfTen * 95000) + (units * 10000))
	return totalCost
}
