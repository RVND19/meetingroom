package entities

import "time"

type Reservation struct {
	Id int `json:"id" gorm:"id,primaryKey"`
	RoomId int `json:"roomId" gorm:"room_id"`
	Room *Room `json:"room" gorm:"foreignKey:RoomId"`
	UserId int `json:"userId" gorm:"user_id"`
	User *User `json:"user" gorm:"foreignKey:UserId"`
	StartDate time.Time `json:"startDate" gorm:"start_date"`
	EndDate time.Time `json:"endDate" gorm:"end_date"`
}

type CreateReservationDto struct {
	RoomId int `json:"roomId"`
	StartDate string `json:"startDate"`
	EndDate string`json:"endDate"`
	Email string `json:"email"`
	Password string `json:"password"`
	
}