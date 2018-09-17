package service_test

import (
	"github.com/Tweeter/src/service"
	"testing"
)

func TestNewUserIsRegistered(t *testing.T) {
	userManager := service.NewUserManager()

	userManager.NewUser("Alvaro", "machicadoa@gmail.com", "Al", "1234")

	users := userManager.Users

	// Validation
	if len(users) != 1 {
		t.Errorf("Expected count 1 but was %d", len(users))
	}

	if users[0].Nick != "Al" {
		t.Errorf("Expected 'Al' but was '%s'", users[0].Nick)
	}
}