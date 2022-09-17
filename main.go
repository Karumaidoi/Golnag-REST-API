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

type Animals struct {
	ID    string `"json:id"`
	Name  string `"json:name"`
	Image string `"json:image"`
	About string `"json:about"`
}

var animal []Animals

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(animal)
}

func createAnimal(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request  being sent 😁")
	// Setting up the headers
	w.Header().Set("Content-Type", "application/json")

	// Instantiating the animal type
	var anim Animals

	// Decoding the information ready to be sent
_:
	json.NewDecoder(r.Body).Decode(&anim)
	anim.ID = strconv.Itoa(rand.Intn(1000000000000000000))
	animal = append(animal, anim)
	// Encoding the RESPONSE received after sending the request
	json.NewEncoder(w).Encode(&anim)
}
func deleteAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range animal {
		if item.ID == param["id"] {
			animal = append(animal[:index], animal[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(animal)
}
func getOneAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	parm := mux.Vars(r)

	for _, item := range animal {
		if item.ID == parm["id"] {
			json.NewEncoder(w).Encode(item)
			break
		}
	}
}

func updateAnimal(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)

	for index, item := range animal {
		if item.ID == param["id"] {
			animal = append(animal[:index], animal[index+1:]...)
			var anim Animals

		_:
			json.NewDecoder(r.Body).Decode(&anim)
			anim.ID = strconv.Itoa(rand.Intn(100000))
			animal = append(animal, anim)
			// Encoding the RESPONSE received after sending the request
			json.NewEncoder(w).Encode(&anim)
		}

	}
	json.NewEncoder(w).Encode(animal)
}

func main() {
	fmt.Println("I love go lang 😂")

	//Connecting with Mongo

	//SEETING UP THE ROUTE
	router := mux.NewRouter()

	//Array of Animals
	animal = append(animal, Animals{ID: strconv.Itoa(rand.Intn(100000)), Name: "Lion", Image: "https://cdn.britannica.com/55/2155-050-604F5A4A/lion.jpg", About: "The lion is a large cat of the genus Panthera native to Africa and India. It has a muscular, broad-chested body, short, rounded head, round ears, and a hairy tuft at the end of its tail. It is sexually dimorphic; adult male lions are larger than females and have a prominent mane"})
	animal = append(animal, Animals{ID: strconv.Itoa(rand.Intn(100000)), Name: "Hyena", Image: "https://cdn.britannica.com/28/130028-050-4CDC40C6/Hyena.jpg", About: "hyena, (family Hyaenidae), any of three species of coarse-furred, doglike carnivores found in Asia and Africa and noted for their scavenging habits. Hyenas have long forelegs and a powerful neck and shoulders for dismembering and carrying prey. "})
	//Make API Routes
	router.HandleFunc("/api/v1/animals", getBooks).Methods("GET")
	router.HandleFunc("/api/v1/animals/create", createAnimal).Methods("POST")
	router.HandleFunc("/api/v1/animals/delete/{id}", deleteAnimal).Methods("DELETE")
	router.HandleFunc("/api/v1/animals/update/{id}", updateAnimal).Methods("PUT")
	router.HandleFunc("/api/v1/animals/{id}", getOneAnimal).Methods("GET")

	//Start our server
	log.Fatal(http.ListenAndServe(":8000", router))
}
