package passenger

import "geektrust/station"

type Passenger struct {
	station.Station
	name string
}

var passengerInstances = make(map[string]*Passenger)

func (self *Passenger) IncrementValueCentral() {
	self.Station.Central++
}

func (self *Passenger) IncrementValueAirport() {
	self.Station.Airport++
}

func (self *Passenger) CurrentValueCentral() int {

	return self.Station.Central

}

func (self *Passenger) CurrentValueAirport() int {

	return self.Station.Airport

}

func (self *Passenger) TotalPassengers() int {
	return self.Station.Airport + self.Station.Central
}

func (self *Passenger) GetPassengerType() string {
	return self.name
}

func NewPassenger(metroNo string, name string) *Passenger {
	if value, ok := passengerInstances[metroNo]; ok {
		return value
	}
	passengerInstances[metroNo] = &Passenger{name: name}
	return passengerInstances[metroNo]
}
