package controllers

import (
	"ProyectoUPCAproval/database"
	"ProyectoUPCAproval/entities"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetProyects(w http.ResponseWriter, r *http.Request) {
	var proyects []entities.Proyecto
	log.Println(">>>>>>>>>GetProyects...")
	database.Instance.Find(&proyects)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(proyects)
}

func CreateProyect(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>>>>>>>CreateProyect...")
	w.Header().Set("Content-Type", "application/json")
	var proyect entities.Proyecto
	err := json.NewDecoder(r.Body).Decode(&proyect)
	if err != nil {
		log.Println(">>>>>>>>>CreateProyect..." + err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "proyect: %+v", proyect)

	database.Instance.Create(&proyect)
	json.NewEncoder(w).Encode(proyect)
}

func GetProyectById(w http.ResponseWriter, r *http.Request) {
	proyectId := mux.Vars(r)["id"]
	if checkIfProyectExists(proyectId) == false {
		json.NewEncoder(w).Encode("proyect Not Found!")
		return
	}
	var proyect entities.Proyecto
	database.Instance.First(&proyect, proyectId)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proyect)
}

func UpdateProyect(w http.ResponseWriter, r *http.Request) {
	proyectId := mux.Vars(r)["id"]
	if checkIfProyectExists(proyectId) == false {
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var proyect entities.Proyecto
	database.Instance.First(&proyect, proyectId)
	json.NewDecoder(r.Body).Decode(&proyect)
	database.Instance.Save(&proyect)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(proyect)
}

func DeleteProyect(w http.ResponseWriter, r *http.Request) {
	log.Println(">>>>>>>>>DeleteProduct...")
	w.Header().Set("Content-Type", "application/json")
	proyectId := mux.Vars(r)["id"]
	if checkIfProyectExists(proyectId) == false {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Product Not Found!")
		return
	}
	var proyect entities.Proyecto
	database.Instance.Delete(&proyect, proyectId)
	json.NewEncoder(w).Encode("Product Deleted Successfully!")
}

func checkIfProyectExists(proyectId string) bool {
	var proyect entities.Proyecto
	database.Instance.First(&proyect, proyectId)
	if proyect.IdProyecto == 0 {
		return false
	}
	return true
}
