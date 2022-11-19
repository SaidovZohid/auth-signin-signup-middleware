package utils

import (
	"log"
	"testing"
)

func TestPassword(t *testing.T) {
	passwords := []string{
		"123456789",
		"987654321",
		"zohid2004",
		"zufar2000",
	}
	for _, password := range passwords {
		hashedPassword, err := HashPassword(password)
		if err != nil {
			log.Fatalf("Failed Test: %v", password)
		}
		err = CheckPassword(password, hashedPassword)
		if err != nil {
			log.Fatalf("Failed Test: %v", password)
		}
		log.Print("Test Passed")
	} 
}
