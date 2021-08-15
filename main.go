package main

import (
	"AwesomeGo/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func addPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "error occured")
		return
	}

	err = models.AddPerson(person.Firstname, person.Lastname)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error occured")
		return
	}

	fmt.Fprint(w, "success")
}

func getAllPerson(w http.ResponseWriter, r *http.Request) {
	persons, err := models.AllPerson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error occured")
		return
	}

	json.NewEncoder(w).Encode(persons)
}

func getPersonById(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "error occured")
		return
	}

	person, err2 := models.PersonById(id)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error occured")
		return
	}

	json.NewEncoder(w).Encode(person)
}

func main() {
	defer models.Database.Close()

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	r.HandleFunc("/person", addPerson).Methods(http.MethodPost)
	r.HandleFunc("/person", getAllPerson).Methods(http.MethodGet)
	r.HandleFunc("/person/{id}", getPersonById).Methods(http.MethodGet)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
