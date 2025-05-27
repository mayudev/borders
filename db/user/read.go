package user

import (
	"log"

	"github.com/melsincostan/argon2id"
	"github.com/melsincostan/borders/db/models"
	"gorm.io/gorm"
)

const (
	otherHash = "$argon2id$v=19$m=65536,t=1,p=4$oen20GX2B0xp6DzIgC4n1A$V+LPf+XkzIcV79pyqlWkIRbCm+KEofGkYAe8AHBJK3A" // hash of bwaa
)

func Authenticate(db *gorm.DB, username, password string) (u *models.User, e error) {
	var user models.User

	if err := db.Model(&models.User{}).Where("name = ?", username).First(&user).Error; err != nil {
		h, e := argon2id.Parse(otherHash)
		if e != nil {
			log.Printf("WARNING: error parsing the timing attack protection hash! This might allow users to be enumerated.")
			return nil, err
		}
		h.Compare("bwaa")
		return nil, err
	}
	return &user, user.Check(password)
}
