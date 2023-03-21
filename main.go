package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Vehicle struct {
	id    int
	Make  string
	Model string
	price int
}

var vehicles = []Vehicle{
	{1, "Toyota", "camry", 50000},
	{2, "Honda", "civic", 90000},
	{3, "Range rover", "vela", 120000},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

func returnCarByBrand(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	carM := vars["make"]
	cars := &[]Vehicle{}
	for _, car := range vehicles {
		if car.Make == carM {
			*cars = append(*cars, car)
		}
	}
	json.NewEncoder(w).Encode(cars)
}

func returnCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert strings")
	}
	for _, car := range vehicles {
		if car.id == carId {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(car)
		}
	}
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert strings")
	}
	var updateCar Vehicle
	json.NewDecoder(r.Body).Decode(&updateCar)
	for k, v := range vehicles {
		if v.id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
			vehicles = append(vehicles, updateCar)
		}
	}
	json.NewEncoder(w).Encode(vehicles)
	w.WriteHeader(http.StatusOK)
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var newCar Vehicle
	json.NewDecoder(r.Body).Decode(&newCar)
	vehicles = append(vehicles, newCar)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)

}

func deleteCarById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	carId, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println("Unable to convert strings")
	}
	for k, v := range vehicles {
		if v.id == carId {
			vehicles = append(vehicles[:k], vehicles[k+1:]...)
		}
		json.NewEncoder(w).Encode(vehicles)
	}
	w.WriteHeader(http.StatusOK)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/cars/", returnAllCars).Methods("GET")
	router.HandleFunc("/cars/make/{make}", returnCarByBrand).Methods("GET")
	router.HandleFunc("/cars/{id}", returnCarById).Methods("GET")
	router.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	router.HandleFunc("/cars", createCar).Methods("POST")
	router.HandleFunc("/cars/{id}", deleteCarById).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
