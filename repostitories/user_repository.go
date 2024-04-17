package repostitories

import (
	"github.com/RVND19/meetingroom/entities"
	"github.com/RVND19/meetingroom/interfaces"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}


func NewUserRepository(db *gorm.DB) interfaces.IUserRepo {
	return &UserRepository{db: db}
}
// CreateUser implements interfaces.IUserRepo.
func (repo *UserRepository) CreateUser(user *entities.User) error {
	return repo.db.Create(user).Error
}

// DeleteUser implements interfaces.IUserRepo.
func (repo *UserRepository) DeleteUser(id int) error {
	return repo.db.Delete(&entities.User{}, id).Error
}

// FindUserByEmail implements interfaces.IUserRepo.
func (repo *UserRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	return &user, repo.db.Where("email = ?", email).First(&user).Error
}

// GetById implements interfaces.IUserRepo.
func (repo *UserRepository) GetById(id int) (*entities.User, error) {
	var user entities.User
	return &user, repo.db.First(&user, id).Error
}

// UpdateUser implements interfaces.IUserRepo.
func (repo *UserRepository) UpdateUser(user *entities.User) error {
	return repo.db.Save(user).Error
}

