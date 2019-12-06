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
