package repositories

import (
	"go.com/railwayticket/database"
	"go.com/railwayticket/models"
)

func GetAllTickets() ([]models.Passenger, error) {
	var passengers []models.Passenger
	result := database.DB.Find(&passengers)
	return passengers, result.Error
}

func GetPassengerByID(id uint) (models.Passenger, error) {
	var passenger models.Passenger
	result := database.DB.First(&passenger, id)
	return passenger, result.Error
}

func SavePassenger(passenger *models.Passenger) error {
	result := database.DB.Save(passenger)
	return result.Error
}

// func DeletePassenger(id uint) error {
// 	result := database.DB.Delete(&models.Passenger{}, id)
// 	return result.Error
// }
