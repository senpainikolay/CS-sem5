package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/senpainikolay/CS-sem5/controllers"
	"github.com/senpainikolay/CS-sem5/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	AuthController controllers.AuthController
)

func init() {
	var err error
	DB, err = gorm.Open(sqlite.Open("golang.db"), &gorm.Config{})
	DB.AutoMigrate(&models.User{})

	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")

	AuthController = controllers.NewAuthController(DB)
}

func RunServer() {
	r := mux.NewRouter()
	r.HandleFunc("/register", RegisterUser).Methods("POST")
	r.HandleFunc("/login", LogInUser).Methods("POST")
	log.Println("Runining on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var usrReg models.RegisterUserInput
	err := json.NewDecoder(r.Body).Decode(&usrReg)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.SignUpUser(usrReg))

}
func LogInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var usrLogIn models.LoginUserInput
	err := json.NewDecoder(r.Body).Decode(&usrLogIn)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	fmt.Fprint(w, AuthController.LoginUser(usrLogIn))

}
