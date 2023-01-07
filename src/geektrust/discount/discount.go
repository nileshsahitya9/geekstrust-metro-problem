package discount

import "geektrust/station"

type Discount struct {
	station.Station
}

var discountInstances = make(map[string]*Discount)

func (self *Discount) IncrementValueCentral(value int) {

	self.Station.Central += value

}

func (self *Discount) IncrementValueAirport(value int) {
	self.Station.Airport += value
}

func (self *Discount) CurrentValueCentral() int {

	return self.Station.Central

}

func (self *Discount) CurrentValueAirport() int {

	return self.Station.Airport

}

func NewDiscount(name string) *Discount {
	if value, ok := discountInstances[name]; ok {
		return value
	}
	discountInstances[name] = &Discount{}
	return discountInstances[name]
}
