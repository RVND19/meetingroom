package entities

type Room struct {
	Id int `json:"id" gorm:"id,primaryKey"`
	Name string `json:"name" gorm:"name"`
}

type CreateRoomDto struct {
	Name string `json:"name"`
}