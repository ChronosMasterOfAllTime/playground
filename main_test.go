package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORMJoin(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	pet := Pet{UserID: &user.ID, Name: "leeroy jenkins"}
	DB.Create(&pet)

	var petResult Pet
	if err := DB.First(&petResult, pet.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}

	var finalResult []UserPet
	DB.Select("users.*, p.name as pet_name").
		Joins("JOIN pets p ON p.user_id = ?", user.ID).
		Find(&finalResult)

	fmt.Println(finalResult)
}

func TestGormInteractions(t *testing.T) {
	// Create a basic table using gorm

	type MyModel struct {
		ID        uint
		VarChar   *time.Time
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt
	}

	// Add one row to the table with a valid date string
	validDateString := "2022-01-01"
	validDate, err := time.Parse("2006-01-02", validDateString)
	assert.NoError(t, err)

	err = DB.Create(&MyModel{VarChar: &validDate}).Error
	if assert.NoError(t, err) {

		// Read the value from the table
		var result MyModel
		err = DB.First(&result).Error
		if assert.NoError(t, err) {

			// Validate that the column can be read
			assert.NotNil(t, result.VarChar)
			assert.Equal(t, validDate, *result.VarChar)
		}
	}
}
