package main

import (
	db "gameback_v1/db"
	"gameback_v1/types"
	utils "gameback_v1/utils"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var s db.Store
var cnf types.Config

func main() {

	cnf = utils.ReadConfig("config.json")

	s.NewStore(cnf.SQLConnection)

	r := mux.NewRouter()

	// TODO
	//r.Use(loggingMiddleware)
	//r.Use(loggingMiddleware)
	//r.Use(loggingMiddleware)

	apiSubV1 := r.PathPrefix("/api/v1").Subrouter()

	apiSubV1.HandleFunc("/categories/", GetCategories).Methods("GET")
	apiSubV1.HandleFunc("/getcategoryitemsbyid/{id}/", GetCategoryItemsById).Methods("GET")
	apiSubV1.HandleFunc("/getitembyid/{id}/", GetItemById).Methods("GET")
	apiSubV1.HandleFunc("/getlocatonbyid/{id}/", GetLocatonById).Methods("GET")

	r.PathPrefix("/pub/img/{*.jpg}").Handler(http.StripPrefix("/pub/img/", http.FileServer(http.Dir(cnf.ImagePath))))
	r.PathPrefix("/pub/loc/{*.xml}").Handler(http.StripPrefix("/pub/loc/", http.FileServer(http.Dir(cnf.LocoFilePath))))

	http.Handle("/", r)

	srv := &http.Server{
		Addr: "0.0.0.0" + cnf.Port,
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	//???
	//go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	log.Println(srv.ListenAndServe())
	//log.Println(http.ListenAndServeTLS(cnf.Port, "s.crt", "s.key", r))

}
