package entities

type User struct {
	Id int `json:"id" gorm:"id,primaryKey"`
	Name string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
	Salt string `json:"salt" gorm:"salt"`
	IsVerified bool `json:"isVerified" gorm:"is_verified"`
	IsAdmin bool `json:"isAdmin" gorm:"is_admin"`
}


type CreateUserDto struct {
	Name string `json:"name" gorm:"name"`
	Email string `json:"email" gorm:"email"`
	Password string `json:"password" gorm:"password"`
}