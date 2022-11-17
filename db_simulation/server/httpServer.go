package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	clientSimulation "github.com/senpainikolay/CS-sem5/db_simulation/client"
)

var client = clientSimulation.GetClientInterfaceSimulation()

func RunDBSimulationServer() {
	r := mux.NewRouter()
	r.HandleFunc("/register/{usr}/{val}", RegisterUser).Methods("POST")
	r.HandleFunc("/login/{usr}/{val}", LogInUser).Methods("POST")
	r.HandleFunc("/delete/{usr}/{val}", DeleteUser).Methods("DELETE")
	log.Println("Runining on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr := vars["usr"]
	val := vars["val"]
	var resp string
	err := client.RegisterCredentials(usr, val)
	if err {
		resp = "Not succesful Registration"
	} else {
		resp = "Succesful registration"
	}
	fmt.Fprint(w, resp)

}
func LogInUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr := vars["usr"]
	val := vars["val"]
	var resp string
	err := client.LogInCredentials(usr, val)
	if err {
		resp = "Not succesful Log In"
	} else {
		resp = "Succesful log in "
	}
	fmt.Fprint(w, resp)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr := vars["usr"]
	val := vars["val"]
	var resp string
	err := client.DeleteCredentials(usr, val)
	if err {
		resp = "Not succesful Delete"
	} else {
		resp = "Succesful Delete"
	}
	fmt.Fprint(w, resp)

}
