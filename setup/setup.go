package setup

import (
	"errors"
	"fmt"

	"github.com/melsincostan/borders/db/country"
	"github.com/melsincostan/borders/db/models"
	"github.com/melsincostan/borders/db/transport"
	"github.com/melsincostan/borders/db/user"
	"gorm.io/gorm"
)

func Setup(db *gorm.DB, username, password string) (err error) {

	if err := db.AutoMigrate(&models.Country{}, &models.Crossing{}, &models.Transport{}, &models.User{}, &models.Setup{}); err != nil {
		return fmt.Errorf("couldn't automigrate tables: %w", err)
	}

	if err := IsSetup(db); !errors.Is(err, ErrNotSetup) {
		return err
	} else if err == nil {
		return ErrAlreadySetup
	}

	for _, table := range []string{"users", "transports", "countries", "crossings", "setups"} {
		if err := db.Exec(fmt.Sprintf("DELETE FROM %s", table)).Error; err != nil {
			return fmt.Errorf("error clearing table '%s': %w", table, err)
		}
	}

	if _, err := user.Create(db, models.UserBase{
		Name: username,
	}, password); err != nil {
		return fmt.Errorf("could not create user: %w", err)
	}

	if _, err := country.CreateBatch(db, country.Default); err != nil {
		return fmt.Errorf("could not add default countries: %w", err)
	}

	if _, err := transport.CreateBatch(db, transport.Default); err != nil {
		return fmt.Errorf("could not add default transports: %w", err)
	}

	setup := models.Setup{
		Version: version,
	}

	if err := db.Model(&models.Setup{}).Create(&setup).Error; err != nil {
		return fmt.Errorf("couldn't mark DB as set up: %w", err)
	}
	return
}
