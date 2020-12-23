package main

func CalculateFuel(moduleMass int) int {
	return moduleMass/3 - 2
}

func SumOfFuels(moduleMasses []int) (fuel int) {
	for _, moduleMass := range moduleMasses {
		fuel += CalculateFuel(moduleMass)
	}
	return
}
