package types

type MetroMethods interface {
	incrementValueCentral(value int)
	incrementValueAirport(value int)
	currentValueCentral() int
	currentValueAirport() int
}
