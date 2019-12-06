package day1

type FuelCounter struct {
	count int
}

func (fc *FuelCounter) Add(i int) {
	fc.count += i
}

func getFuelConsumption(mass int) int {
	return mass/3 - 2
}

func fuelOfFuel(mass int) int {
	result := 0
	for {
		mass = getFuelConsumption(mass)

		if mass > 0 {
			result += mass
		} else {
			break
		}
	}
	return result
}
