package migration

import (
	"fmt"
	"log"

	"ksni.com/crud/database"
	"ksni.com/crud/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{})

	if (err != nil) {
		log.Println(err)
	}

	fmt.Println("Database migrated")
}