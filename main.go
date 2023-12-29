package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Car struct {
	ID      string   `json:"id"`
	Model   string   `json:"model"`
	Company *Company `json:"company"`
}

type Company struct {
	Name string `json:"name"`
}

var cars []Car

func listAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}

func deleteOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func listOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range cars {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func addCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var car Car

	_ = json.NewDecoder(r.Body).Decode(&car)
	car.ID = strconv.Itoa(rand.Intn(10000000))
	cars = append(cars, car)
	json.NewEncoder(w).Encode(car)
}

func updateOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range cars {
		if item.ID == params["id"] {
			cars = append(cars[:index], cars[index+1:]...)
			var car Car
			_ = json.NewDecoder(r.Body).Decode(&car)
			car.ID = params["id"]
			cars = append(cars, car)
			json.NewEncoder(w).Encode(car)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	cars = append(cars, Car{ID: "1", Model: "C200", Company: &Company{Name: "Mercedes Benz"}})
	cars = append(cars, Car{ID: "2", Model: "Accord", Company: &Company{Name: "Honda"}})
	cars = append(cars, Car{ID: "3", Model: "Avensis", Company: &Company{Name: "Toyota"}})
	cars = append(cars, Car{ID: "4", Model: "Corolla", Company: &Company{Name: "Toyota"}})
	cars = append(cars, Car{ID: "5", Model: "Yaris", Company: &Company{Name: "Toyota"}})
	cars = append(cars, Car{ID: "6", Model: "Prado", Company: &Company{Name: "Toyota"}})
	cars = append(cars, Car{ID: "7", Model: "Highlander", Company: &Company{Name: "Toyota"}})
	cars = append(cars, Car{ID: "8", Model: "718 Cayman", Company: &Company{Name: "Porsche"}})
	cars = append(cars, Car{ID: "9", Model: "911", Company: &Company{Name: "Porsche"}})
	cars = append(cars, Car{ID: "10", Model: "Cayenne", Company: &Company{Name: "Porsche"}})
	cars = append(cars, Car{ID: "11", Model: "Boxster", Company: &Company{Name: "Porsche"}})
	cars = append(cars, Car{ID: "12", Model: "Carrera GT", Company: &Company{Name: "Porsche"}})
	cars = append(cars, Car{ID: "13", Model: "1500 Extended Cab", Company: &Company{Name: "Chevrolet"}})
	cars = append(cars, Car{ID: "14", Model: "2500 Cre Cab", Company: &Company{Name: "Chevrolet"}})
	cars = append(cars, Car{ID: "15", Model: "Equinox", Company: &Company{Name: "Chevrolet"}})
	cars = append(cars, Car{ID: "16", Model: "Express 1500 Cargo", Company: &Company{Name: "Chevrolet"}})
	cars = append(cars, Car{ID: "17", Model: "SLS-Class", Company: &Company{Name: "Mercedes Benz"}})
	cars = append(cars, Car{ID: "18", Model: "Sprinter 1500 Crew", Company: &Company{Name: "Mercedes Benz"}})
	cars = append(cars, Car{ID: "19", Model: "Charger", Company: &Company{Name: "Dodge"}})
	cars = append(cars, Car{ID: "20", Model: "Colt", Company: &Company{Name: "Dodge"}})
	cars = append(cars, Car{ID: "21", Model: "Avenger", Company: &Company{Name: "Dodge"}})
	cars = append(cars, Car{ID: "22", Model: "Caliber", Company: &Company{Name: "Dodge"}})

	r.HandleFunc("/cars", listAll).Methods("GET")
	r.HandleFunc("/cars/{id}", listOne).Methods("GET")
	r.HandleFunc("/cars", addCar).Methods("POST")
	r.HandleFunc("/cars/{id}", updateOne).Methods("PUT")
	r.HandleFunc("/cars/{id}", deleteOne).Methods("DELETE")

	fmt.Printf("Server at port http://localhost:8081\n")
	log.Fatal(http.ListenAndServe(":8081", r))
}
