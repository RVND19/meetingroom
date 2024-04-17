package interfaces

import (
	"time"

	"github.com/RVND19/meetingroom/entities"
)

type IUserRepo interface {

	CreateUser(user *entities.User) error
	GetById(id int) (*entities.User, error)
	FindUserByEmail(email string) (*entities.User, error)
	UpdateUser(user *entities.User) error
	DeleteUser(id int) error
}

type IRoomRepo interface {
	CreateRoom(room *entities.Room) error
	GetById(id int) (*entities.Room, error)
	UpdateRoom(room *entities.Room) error
	DeleteRoom(id int) error
}


type IReservationRepo interface {
	CreateReservation(reservation *entities.Reservation) error
	GetById(id int) (*entities.Reservation, error)
	UpdateReservation(reservation *entities.Reservation) error
	DeleteReservation(id int) error
	GetAvailableRoom(startDate time.Time, endDate time.Time) ([]*entities.Room, error)
	IsRoomAvailable(roomId int, startDate time.Time, endDate time.Time) (bool, error)
}