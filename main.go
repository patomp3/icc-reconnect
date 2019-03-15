package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type appConfig struct {
	port          string
	reconnecturl  string
	cancelurl     string
	disconnecturl string
}

var cfg appConfig

func main() {
	log.Printf("##### ICC Collection Service Started #####")

	// For no assign parameter env. using default to Test
	var env string
	if len(os.Args) > 1 {
		env = strings.ToLower(os.Args[1])
	} else {
		env = "development"
	}

	// Load configuration
	viper.SetConfigName("app")    // no need to include file extension
	viper.AddConfigPath("config") // set the path of your config file
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found..." + err.Error())
	} else {
		// read config file
		cfg.port = viper.GetString(env + ".port")
		cfg.reconnecturl = viper.GetString(env + ".reconnecturl")
		cfg.disconnecturl = viper.GetString(env + ".disconnecturl")
		cfg.cancelurl = viper.GetString(env + ".cancelurl")

		log.Printf("## Loading Configuration")
		log.Printf("## Env\t= %s", env)
		log.Printf("## Port\t= %s", cfg.port)
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

	log.Printf("## Reconnect Request incoming...")
	//log.Printf("## >> Customer: %d", req.Customer.CustomerID)
	log.Printf("## >> %v", req)

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

	log.Printf("## Disconnect Request incoming...")
	//log.Printf("## >> Customer: %d", req.Customer.CustomerID)
	log.Printf("## >> %v", req)

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

	log.Printf("## Cancel Request incoming...")
	//log.Printf("## >> Customer: %d", req.Customer.CustomerID)
	log.Printf("## >> %v", req)

	//call recon api
	var res CancelResult
	res = CancelProduct(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
