package repository

import (
	"errors"
	"log"

	"github.com/jchukwum/ai-fitted-user/pkg/cfg"
	"github.com/jchukwum/ai-fitted-user/pkg/model"
)

// Init initialises the database with the objects
func Init() {
	err := cfg.Db.AutoMigrate(&model.Profile{})
	if err != nil {
		log.Println(err)
	}
}

// EmailExists checks if there's a user with same email
func EmailExists(email string) bool {
	resp := cfg.Db.First(&email, "email = ?", email)
	return resp.RowsAffected != 0
}

// SaveProfile to save profile
func SaveProfile(p *model.Profile) {
	var old model.Profile
	resp := cfg.Db.First(&old, "user_id = ?", p.UserID) // Check if record exists
	if resp.Error != nil {
		cfg.Db.Create(p)
	} else {
		cfg.Db.Save(p)
	}
}

// GetProfile to save user
func GetProfile(u string) (*model.Profile, error) {
	var p model.Profile
	resp := cfg.Db.Where("user_id = ?", u).First(&p)
	if resp.RowsAffected == 0 {
		return nil, errors.New("Couldn't find profile in database")
	}
	return &p, nil
}
