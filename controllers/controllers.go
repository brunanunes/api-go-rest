package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/brunanunes/go-rest-api/database"
	"github.com/brunanunes/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func Personalities(w http.ResponseWriter, r *http.Request) {
	var personality []models.Personality
	database.DB.Find(&personality)
	json.NewEncoder(w).Encode(personality)
}

func Personality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)

	database.DB.Create(&personality)
	json.NewEncoder(w).Encode(personality)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.Delete(&personality, id)
	json.NewEncoder(w).Encode(personality)
}

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	database.DB.First(&personality, id)
	json.NewDecoder(r.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(w).Encode(personality)
}
