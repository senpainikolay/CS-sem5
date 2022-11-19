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
	r.HandleFunc("/login/{usr}/{val}/{token}", LogInUser).Methods("POST")
	log.Println("Runining on localhost:8080")
	http.ListenAndServe(":8080", r)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr := vars["usr"]
	val := vars["val"]
	token := clientSimulation.RegisterInterface(client, usr, val) //  client.RegisterCredentials(usr, val)
	fmt.Fprint(w, token)

}
func LogInUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	usr := vars["usr"]
	val := vars["val"]
	token := vars["token"]
	resp := clientSimulation.LogInInterface(client, usr, val, token) // client.LogInCredentials(usr, val, token)
	fmt.Fprint(w, resp)

}
