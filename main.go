package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type config struct {
	port          string
	reconnecturl  string
	cancelurl     string
	disconnecturl string
}

func main() {
	fmt.Println("Start collection service.....")

	//env := os.Args[1]
	var env string
	var cfg config

	// For no assign parameter env. using default to Test
	if len(os.Args) > 1 {
		env = os.Args[1]
	} else {
		env = "Test"
	}

	viper.SetConfigName("app")    // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found..." + err.Error())
	} else {
		if env == "production" {
			cfg.port = viper.GetString("production.port")
			cfg.reconnecturl = viper.GetString("production.reconnecturl")
			cfg.disconnecturl = viper.GetString("production.disconnecturl")
			cfg.cancelurl = viper.GetString("production.cancelurl")
		} else {
			cfg.port = viper.GetString("development.port")
			cfg.reconnecturl = viper.GetString("development.reconnecturl")
			cfg.disconnecturl = viper.GetString("development.disconnecturl")
			cfg.cancelurl = viper.GetString("development.cancelurl")
		}
		fmt.Println("Env=" + env)
		fmt.Println("serverPort=" + cfg.port)
		fmt.Println("reconnecturl=" + cfg.reconnecturl)
		fmt.Println("disconnecturl=" + cfg.disconnecturl)
		fmt.Println("cancelurl=" + cfg.cancelurl)
	}

	router := mux.NewRouter()
	router.HandleFunc("/reconnect", reconnect).Methods("POST")
	router.HandleFunc("/disconnect", disconnect).Methods("POST")
	router.HandleFunc("/cancel", cancel).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+cfg.port, router))
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

	log.Println("Reconnect")
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

	log.Println("Disconnect")
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

	log.Println("Cancel")
	log.Println(req)

	//call recon api
	var res CancelResult
	res = CancelProduct(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
