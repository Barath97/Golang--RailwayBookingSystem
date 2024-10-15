package services

import (
	"errors"

	"go.com/railwayticket/models"
	"go.com/railwayticket/repositories"
)

var (
	totalLowerBerths  = 10
	totalMiddleBerths = 10
	totalUpperBerths  = 10
	totalRacTickets   = 5

	lowerBerthPositions  []int
	middleBerthPositions []int
	upperBerthPositions  []int
	racPositions         []int
)

func init() {
	//Dynamically generate the berth positions
	lowerBerthPositions = make([]int, totalLowerBerths)
	middleBerthPositions = make([]int, totalMiddleBerths)
	upperBerthPositions = make([]int, totalUpperBerths)
	racPositions = make([]int, totalRacTickets)

	for i := 0; i < totalLowerBerths; i++ {
		lowerBerthPositions[i] = i + 1
	}

	for i := 0; i < totalMiddleBerths; i++ {
		middleBerthPositions[i] = i + 1
	}

	for i := 0; i < totalUpperBerths; i++ {
		upperBerthPositions[i] = i + 1
	}

	for i := 0; i < totalRacTickets; i++ {
		racPositions[i] = i + 1
	}

}

func BookTicket(passenger *models.Passenger) (bool, error) {

	if totalLowerBerths > 0 || totalMiddleBerths > 0 || totalUpperBerths > 0 || totalRacTickets > 0 {
		switch passenger.BerthPreference {
		case "L":
			if totalLowerBerths > 0 {
				seatAllocation("L", passenger)
				totalLowerBerths--
				lowerBerthPositions = lowerBerthPositions[1:]
				repositories.SavePassenger(passenger)
				return true, nil
			}
		case "M":
			if totalMiddleBerths > 0 {
				seatAllocation("M", passenger)
				totalMiddleBerths--
				upperBerthPositions = upperBerthPositions[1:]
				repositories.SavePassenger(passenger)
				return true, nil
			}
		case "U":
			if totalUpperBerths > 0 {
				seatAllocation("U", passenger)
				totalUpperBerths--
				upperBerthPositions = upperBerthPositions[1:]
				repositories.SavePassenger(passenger)
				return true, nil
			}
		}

		if totalLowerBerths > 0 {
			seatAllocation("L", passenger)
			totalLowerBerths--
			lowerBerthPositions = lowerBerthPositions[1:]
			repositories.SavePassenger(passenger)
			return true, nil
		}

		if totalMiddleBerths > 0 {
			seatAllocation("M", passenger)
			totalMiddleBerths--
			middleBerthPositions = middleBerthPositions[1:]
			repositories.SavePassenger(passenger)
			return true, nil
		}

		if totalUpperBerths > 0 {
			seatAllocation("U", passenger)
			totalUpperBerths--
			middleBerthPositions = middleBerthPositions[1:]
			repositories.SavePassenger(passenger)
			return true, nil
		}

		if totalRacTickets > 0 {
			seatAllocation("RAC", passenger)
			totalRacTickets--
			racPositions = racPositions[1:]
			repositories.SavePassenger(passenger)
			return true, nil
		}
	}
	return false, errors.New("no berth avaiable")
}

func seatAllocation(preference string, passenger *models.Passenger) {
	switch preference {
	case "L":
		seat := lowerBerthPositions[0]
		passenger.SeatNumber = preference + string(rune(seat+'0'))
	case "M":
		seat := middleBerthPositions[0]
		passenger.SeatNumber = preference + string(rune(seat+'0'))
	case "U":
		seat := upperBerthPositions[0]
		passenger.SeatNumber = preference + string(rune(seat+'0'))
	case "RAC":
		seat := racPositions[0]
		passenger.SeatNumber = preference + string(rune(seat+'0'))
	}
}

func GetAllTickets() ([]models.Passenger, error) {
	return repositories.GetAllTickets()
}

func GetTicketDetails(id int) (models.Passenger, error) {
	return repositories.GetPassengerByID(uint(id))
}
