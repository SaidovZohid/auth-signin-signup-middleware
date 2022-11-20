package utils

import (
	"log"
	"testing"
	"time"
)

func TestToken(t *testing.T) {
	token, _, err := CreateToken("zohid", "saidov", "zohidsaidov17@gmail.com", time.Hour*24*30)
	if err != nil {
		log.Fatalln("Failed Test")
	}
	log.Print(token)
}
