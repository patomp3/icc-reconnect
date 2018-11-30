package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Start recon service.....")

	router := mux.NewRouter()
	router.HandleFunc("/reconnect", reconnect).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", router))
}

func reconnect(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//Read Json Request
	var req ReconRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	log.Println(req)

	//call recon api
	var res ReconResult
	res = ReconnectProduct(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
