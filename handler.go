package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func cpuHandler(w http.ResponseWriter, r *http.Request) {
	c, error := getCpuInfos()
	if error != nil {
		log.Panic(error)
	}
	j, error := json.Marshal(c)
	if error != nil {
		log.Panic(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func memHandler(w http.ResponseWriter, r *http.Request) {
	m, error := getMemInfos()
	if error != nil {
		log.Panic(error)
	}
	j, error := json.Marshal(m)
	if error != nil {
		log.Panic(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func diskHandler(w http.ResponseWriter, r *http.Request) {
	d, error := getDiskInfos()
	if error != nil {
		log.Panic(error)
	}
	j, error := json.Marshal(d)
	if error != nil {
		log.Panic(error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

func getRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/cpu", cpuHandler)
	r.HandleFunc("/mem", memHandler)
	r.HandleFunc("/disk", diskHandler)
	return r
}
