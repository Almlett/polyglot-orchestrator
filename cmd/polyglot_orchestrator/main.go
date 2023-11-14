package main

import (
	"PolyglotOrchestrator/config"
	"fmt"

	"PolyglotOrchestrator/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.New()
	dsn := cfg.GetDBURL()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
	}

	err = db.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{}, &models.UserPermission{}, &models.UserRoles{}, &models.RolePermission{}, &models.Token{})
	if err != nil {
		fmt.Println("Error running migrations: ", err)
		return
	}

	fmt.Println("Migration did run successfully")
}
