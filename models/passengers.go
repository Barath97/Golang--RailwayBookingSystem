package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model
	Name            string `json:"name"`
	Age             int    `json:"age"`
	Gender          string `json:"gender"`
	TrainNumber     string `json:"train_number"`
	BerthPreference string `json:"berth_preference"`
	SeatNumber      string `json:"seat_number"`
}

func (p Passenger) TableName() string {
	return "passengers"
}
