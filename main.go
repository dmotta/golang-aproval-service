package main

import (
	"ProyectoUPCAproval/controllers"
	"ProyectoUPCAproval/database"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {

	// Load Configurations from config.json using Viper
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true)

	// Register Routes
	RegisterProductRoutes(router)

	// Start the server
	log.Println(fmt.Sprintf("Starting Server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/proyects", controllers.GetProyects).Methods("GET")
	router.HandleFunc("/api/proyects/{id}", controllers.GetProyectById).Methods("GET")
	router.HandleFunc("/api/proyects", controllers.CreateProyect).Methods("POST")
	router.HandleFunc("/api/proyects/{id}", controllers.UpdateProyect).Methods("PUT")
	router.HandleFunc("/api/proyects/{id}", controllers.DeleteProyect).Methods("DELETE")
}
