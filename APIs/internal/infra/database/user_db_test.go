package database

import (
	"testing"

	"github.com/michaelquesado/baby-step-go/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateUser(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{})

	repo := NewUserRepo(db)
	u, _ := entity.NewUser("any-user", "mail@mail.com", "any-pass")
	err = repo.Create(u)

	assert.Nil(t, err)

	registedUser, err := repo.FindByEmail("mail@mail.com")

	assert.Nil(t, err)
	assert.Equal(t, u.Name, registedUser.Name)
	assert.Equal(t, u.ID, registedUser.ID)
}
