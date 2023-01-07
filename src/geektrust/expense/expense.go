package expense

import "geektrust/station"

type Expense struct {
	station.Station
}

var expenseInstances = make(map[string]*Expense)

func (self *Expense) IncrementValueCentral(value int) {

	self.Station.Central += value

}

func (self *Expense) IncrementValueAirport(value int) {
	self.Station.Airport += value
}

func (self *Expense) CurrentValueCentral() int {

	return self.Station.Central

}

func (self *Expense) CurrentValueAirport() int {

	return self.Station.Airport

}

func NewExpense(name string) *Expense {
	if value, ok := expenseInstances[name]; ok {
		return value
	}
	expenseInstances[name] = &Expense{}
	return expenseInstances[name]
}
