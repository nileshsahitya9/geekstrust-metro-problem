package calculation

import (
	"fmt"
	"geektrust/discount"
	"geektrust/expense"
	"geektrust/passenger"
	"geektrust/variables"
	"sort"
	"strconv"
)

var passengerBalance = make(map[string]int)

var charges = map[string]int{
	variables.Adult:         200,
	variables.Kid:           50,
	variables.SeniorCitizen: 100,
}

func Calculation(userInput [][]string) {

	for _, input := range userInput {
		switch input[0] {
		case variables.Balance:
			balance(input[1], input[2])
		case variables.CheckIn:
			checkINfunc(input[1], input[2], input[3])
		case variables.PrintSummary:
			printSummary()
		}
	}

}

func balance(metroNo string, balance string) {
	value, err := strconv.Atoi(balance)
	if err != nil {
		panic(err)
	}
	passengerBalance[metroNo] = value
}

func checkINfunc(metroNo string, passengerType string, fromStation string) {
	amount := variables.ZeroValue
	value := passengerBalance[metroNo]

	switch fromStation {
	case variables.Central:
		{
			amount = centralProcessing(metroNo, passengerType, fromStation)
		}
	case variables.Airport:
		{
			amount = airportProcessing(metroNo, passengerType, fromStation)
		}
	}
	updateAmount(metroNo, value, amount)
}

func printSummary() {
	centralBilling := variables.ZeroValue
	airportBilling := variables.ZeroValue
	centralDiscount := variables.ZeroValue
	airportDiscount := variables.ZeroValue
	centralPassengers := make(map[string]int)
	airportPassengers := make(map[string]int)

	metroMetrics(&centralBilling, &airportBilling, &centralDiscount, &airportDiscount, centralPassengers, airportPassengers)

	centralList := sortPassenger(centralPassengers)
	airportList := sortPassenger(airportPassengers)

	formattedOutput(variables.Central, centralBilling, centralDiscount, centralList, centralPassengers)
	formattedOutput(variables.Airport, airportBilling, airportDiscount, airportList, airportPassengers)
}

func tripCosting(passengerType string, trip int) (int, bool) {
	amount := charges[passengerType]
	isDiscounted := false
	if trip%variables.DividendNumber == variables.ZeroValue {
		amount = int((float64((trip / variables.DividendNumber) * amount)) * variables.TripDiscount)
		isDiscounted = true
	}

	return amount, isDiscounted
}

func taxCalc(value int, amount int) int {
	tax := variables.ZeroValue
	if value < amount {
		tax = int(float64(amount) * variables.TaxCharge)
	}

	return tax
}

func updateAmount(metroNo string, value int, amount int) {
	if value > variables.ZeroValue {
		remainingAmount := value - amount
		passengerBalance[metroNo] = remainingAmount
		if remainingAmount < variables.ZeroValue {
			passengerBalance[metroNo] = variables.ZeroValue
		}
	}
}

func initializeInstances(metroNo string, passengerType string) (*passenger.Passenger, *discount.Discount, *expense.Expense) {
	return passenger.NewPassenger(metroNo, passengerType), discount.NewDiscount(metroNo), expense.NewExpense(metroNo)
}

func sortPassenger(m map[string]int) []string {
	keys := make([]string, variables.ZeroValue, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})

	return keys
}

func formattedOutput(stationType string, billing int, discount int, passengerList []string, passengerCount map[string]int) {
	fmt.Println("TOTAL_COLLECTION", " ", stationType, " ", billing, " ", discount)
	fmt.Println("PASSENGER_TYPE_SUMMARY")
	for _, element := range passengerList {
		fmt.Println(element, " ", passengerCount[element])
	}
}

func helper(passenger *passenger.Passenger, discount *discount.Discount, passengerType string, fromStation string) int {
	trip := passenger.TotalPassengers()

	amount := charges[passengerType]
	amount, isDiscounted := tripCosting(passengerType, trip)
	if isDiscounted {
		if fromStation == variables.Airport {
			discount.IncrementValueAirport(amount)
		}

		if fromStation == variables.Central {
			discount.IncrementValueCentral(amount)
		}
	}

	return amount
}

func centralProcessing(metroNo string, passengerType string, fromStation string) int {
	passenger, discount, expense := initializeInstances(metroNo, passengerType)
	passenger.IncrementValueCentral()
	amount := 0
	if value, ok := passengerBalance[metroNo]; ok {
		amount = helper(passenger, discount, passengerType, fromStation)
		tax := taxCalc(value, amount-value)
		expense.IncrementValueCentral((amount + tax))

	}
	return amount
}

func airportProcessing(metroNo string, passengerType string, fromStation string) int {
	passenger, discount, expense := initializeInstances(metroNo, passengerType)
	passenger.IncrementValueAirport()
	amount := 0
	if value, ok := passengerBalance[metroNo]; ok {
		amount = helper(passenger, discount, passengerType, fromStation)
		tax := taxCalc(value, amount-value)
		expense.IncrementValueAirport((amount + tax))
	}
	return amount
}

func metroMetrics(centralBilling, airportBilling, centralDiscount, airportDiscount *int, centralPassengers, airportPassengers map[string]int) {
	for key, _ := range passengerBalance {
		passenger, discount, expense := initializeInstances(key, "")
		*centralBilling += expense.CurrentValueCentral()
		*airportBilling += expense.CurrentValueAirport()
		*centralDiscount += discount.CurrentValueCentral()
		*airportDiscount += discount.CurrentValueAirport()

		if passenger.CurrentValueCentral() > variables.ZeroValue {
			centralPassengers[passenger.GetPassengerType()] += passenger.CurrentValueCentral()
		}

		if passenger.CurrentValueAirport() > variables.ZeroValue {
			airportPassengers[passenger.GetPassengerType()] += passenger.CurrentValueAirport()
		}
	}
}
