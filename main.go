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
	router.HandleFunc("/disconnect", disconnect).Methods("POST")
	router.HandleFunc("/cancel", cancel).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
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

func disconnect(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//Read Json Request
	var req DisconRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	log.Println(req)

	//call recon api
	var res DisconResult
	res = DisconnectProduct(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func cancel(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	//Read Json Request
	var req CancelRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		panic(err)
	}

	log.Println(req)

	//call recon api
	var res CancelResult
	res = CancelProduct(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
