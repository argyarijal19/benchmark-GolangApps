package repository

import (
	"context"
	"golang-apps/helper"
	"golang-apps/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	Database *gorm.DB
}

func UserTable(db *gorm.DB) *UserRepository {
	return &UserRepository{Database: db.Table("users").Session(&gorm.Session{}).WithContext(context.Background())}
}

func (ut *UserRepository) GetAllUser() ([]models.UserModel, error) {
	var users []models.UserModel
	results := ut.Database.Find(&users)
	if results.Error != nil {
		return nil, results.Error
	}

	return users, nil
}

func (ut *UserRepository) RegisterUser(user models.UserModel) error {
	hashedID := helper.GenerateUniqueID(user.UserName, user.NamaLengkap)
	hashedPW, err := helper.HashPassword(user.Password)
	if err != nil {
		return err
	}
	dataUser := models.UserModel{
		ID:          hashedID,
		UserName:    user.UserName,
		Password:    hashedPW,
		NamaLengkap: user.NamaLengkap,
	}

	result := ut.Database.Create(&dataUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
