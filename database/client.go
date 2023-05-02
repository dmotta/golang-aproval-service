package database

import (
	"ProyectoUPCAproval/entities"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var err error

func Connect(connectionString string) {
	Instance, err = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database...")
}

func Migrate() {
	Instance.AutoMigrate(&entities.Proyecto{})
	log.Println("Database Migration Completed...")
}
