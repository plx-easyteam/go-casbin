package controllers

import (
	"encoding/json"
	"go-casbin/db"
	"go-casbin/models"
	"log"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	u := db.Instance.Find(&users)

	if u.Error != nil {
		log.Println("::: Error :::", u.Error)

		err := map[string]interface{}{
			"status": 502,
			"message": "Cannot get users",
		}
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)

	log.Println("::: CreatUser :::")

	if err != nil {
		log.Println("::: Error :::", err)

		err := map[string]interface{}{
			"status": 502,
			"message": "Cannot get users",
		}
		json.NewEncoder(w).Encode(err)
		return
	} else {
		db.Instance.Create(&user)
	}

	json.NewEncoder(w).Encode(user)
}