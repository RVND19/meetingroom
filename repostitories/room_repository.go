package repostitories

import (
	"errors"

	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"gorm.io/gorm"
)


type RoomRepository struct {
	db *gorm.DB
}

func NewRoomRepository(db *gorm.DB) interfaces.IRoomRepo {
	return &RoomRepository{db: db}
}

func (repo *RoomRepository) CreateRoom(room *entities.Room) error{
	return repo.db.Create(room).Error
}
func (repo *RoomRepository) GetById(id int) (*entities.Room, error) {
	var room entities.Room
	return &room, repo.db.First(&room, id).Error


}
func (repo *RoomRepository) UpdateRoom(room *entities.Room) error {
	return errors.New("implement me")
}
func (repo *RoomRepository) DeleteRoom(id int) error {
	return errors.New("implement me")
}