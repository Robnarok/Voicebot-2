package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Robnarok/Voicebot-2/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("Hello There!")
	database_Url := os.Getenv("DATABASE_URL")
	if database_Url == "" {
		log.Fatal("Database ENV not Set!")
		os.Exit(1)
	}

	db, err := gorm.Open(postgres.Open(database_Url), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	db.AutoMigrate(&model.Channel{})

	if err != nil {
		log.Fatalf("Error migrating database schema: %v", err)
	}

	newEntry := model.CreateChannel("DEF", "ABC", "GHI")

	//testEntry := model.Channel{VoiceName: "AAA", TextName: "BBB", CategoryName: "CCC"}
	db.Create(&newEntry)
}
