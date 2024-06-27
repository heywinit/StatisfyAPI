package repositories

import (
	"StatisfyAPI/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// Create creates a new user record in the database.
func (repo *UserRepository) Create(user *models.User) error {
	return repo.db.Create(user).Error
}

// FindByID finds a user by ID in the database.
func (repo *UserRepository) FindByID(id uint) (*models.User, error) {
	var user models.User
	err := repo.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByUsername finds a user by username in the database.
func (repo *UserRepository) FindByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Update updates a user by ID in the database.
func (repo *UserRepository) Update(id uint, user *models.User) error {
	return repo.db.Model(&models.User{}).Where("id = ?", id).Updates(user).Error
}

// Delete deletes a user by ID in the database.
func (repo *UserRepository) Delete(id uint) error {
	return repo.db.Delete(&models.User{}, id).Error
}
