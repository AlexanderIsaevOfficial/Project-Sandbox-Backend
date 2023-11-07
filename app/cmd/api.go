package main

import (
	"encoding/json"
	"fmt"
	"gameback_v1/types"
	utils "gameback_v1/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

/*
func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+":3030"+req.URL.String(),
		http.StatusMovedPermanently)
}*/

// Ready
func GetCategories(w http.ResponseWriter, r *http.Request) {

	data, err := s.GetMainCategoryList(11)
	if err != nil {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
		return
	}

	resp, _ := utils.ConvertCategoryDBRequestToMainResponse(data)

	resp.Status = true

	b, err := json.Marshal(resp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
	}

	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(b))
}

// Ready
func GetCategoryItemsById(w http.ResponseWriter, r *http.Request) {

	tn := time.Now()
	defer func() { log.Println("Handler GetCategoryItemsById, time=", time.Since(tn).Microseconds()) }()

	vars := mux.Vars(r)

	urlParams := r.URL.Query().Get("p")
	var bucket int

	if strId, ok := vars["id"]; ok {

		if i, err := strconv.Atoi(urlParams); len(urlParams) > 0 && err == nil {
			bucket = i
		}

		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id dont parse int"}`))
			return
		}

		if id < 1 {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id must be more 0"}`))
			return
		}

		items, err := s.GetCategoryItemsById(id, bucket, 20)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
			return
		}

		if len(items) < 1 {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Empty data"}`))
			return
		} else {
			res := types.ItemsResponse{Status: true, Items: items}

			b, err := json.Marshal(res)
			if err != nil {
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
			}
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(b))
			return
		}

	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"False" , "message":"Uncorrect id"}`))
		return
	}
}

// Ready
func GetItemById(w http.ResponseWriter, r *http.Request) {

	tn := time.Now()
	defer func() { log.Println("Handler GetItemById, time =", time.Since(tn).Microseconds()) }()

	vars := mux.Vars(r)

	if strId, ok := vars["id"]; ok {

		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id dont parse int"}`))
			return
		}

		if id < 1 {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id must be more 0"}`))
			return
		}

		item, err := s.GetItemById(id)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
			return
		}

		b, err := json.Marshal(item)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(b))

	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"False" , "message":"Uncorrect id"}`))
		return
	}
}

// Ready send file xml
func GetLocatonById(w http.ResponseWriter, r *http.Request) {

	tn := time.Now()
	defer func() { log.Println("Handler GetLocatonById, time =", time.Since(tn).Microseconds()) }()

	vars := mux.Vars(r)

	if strId, ok := vars["id"]; ok {

		id, err := strconv.Atoi(strId)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id dont parse int"}`))
			return
		}

		if id < 1 {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(`{"status":"False" , "message":"Id must be more 0"}`))
			return
		}

		loc, err := s.GetLocatonById(id)
		if err != nil {
			w.WriteHeader(200)
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
			return
		}

		b, err := os.ReadFile(cnf.LocoFilePath + loc.FilePath)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(fmt.Sprintf(`{"status":"False" , "message":"%v"}`, err)))
		}

		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		w.Write([]byte(b))

	} else {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"False" , "message":"Uncorrect id"}`))
		return
	}
}
