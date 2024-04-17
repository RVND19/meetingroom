package repostitories

import (
	"errors"
	"fmt"
	"time"

	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"gorm.io/gorm"
)

type ReservationRepository struct {
	db *gorm.DB
}



func NewReservationRepository(db *gorm.DB) interfaces.IReservationRepo {
	return &ReservationRepository{db: db}
}

// CreateReservation implements interfaces.IReservationRepo.
func (repo *ReservationRepository) CreateReservation(reservation *entities.Reservation) error {
	return repo.db.Create(reservation).Error
}

// DeleteReservation implements interfaces.IReservationRepo.
func (repo *ReservationRepository) DeleteReservation(id int) error {
	return repo.db.Delete(&entities.Reservation{}, id).Error
}

// GetById implements interfaces.IReservationRepo.
func (repo *ReservationRepository) GetById(id int) (*entities.Reservation, error) {
	var reservation entities.Reservation
	return &reservation, repo.db.Preload("Room").Preload("User").First(&reservation, id).Error
}

// UpdateReservation implements interfaces.IReservationRepo.
func (repo *ReservationRepository) UpdateReservation(reservation *entities.Reservation) error {
	return repo.db.Save(reservation).Error
}
// GetAvailableRoom implements interfaces.IReservationRepo.
func (repo *ReservationRepository) GetAvailableRoom(startDate time.Time, endDate time.Time) ([]*entities.Room, error) {
	return nil, errors.New("unimplemented")
}

// IsRoomAvailable implements interfaces.IReservationRepo.
func (repo *ReservationRepository) IsRoomAvailable(roomId int, startDate time.Time, endDate time.Time) (bool, error) {
	var reservation []entities.Reservation
	rs:=repo.db.Where("room_id = ? AND "+
					  "((start_date <= ? AND end_date >= ?) OR (start_date <= ? AND end_date >= ?))", roomId, startDate, startDate, endDate, endDate).Find(&reservation)
	if rs.Error != nil  {
		return false, rs.Error
	}
	if len(reservation) == 0 {
		return true, nil
	}else{
		fmt.Printf("booked : %+v\n",reservation)
		return false, nil
	}
}