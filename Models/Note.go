//package Models
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Note struct {
	Id          int
	Title       string
	Descripcion string
	CreatedAt   time.Time
}

var noteStore = make(map[string]Note)
var id int

func getNotesHandler(w http.ResponseWriter, r *http.Request) {
	var notes []Note
	for _, v := range noteStore {
		notes = append(notes, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(notes)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func postNoteHandler(w http.ResponseWriter, r *http.Request) {
	var note Note
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		panic(err)
	}
	note.CreatedAt = time.Now()
	id++
	key := strconv.Itoa(id)
	note.Id = id
	noteStore[key] = note
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(note)
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func putNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var noteUpdate Note
	err := json.NewDecoder(r.Body).Decode(&noteUpdate)
	if err != nil {
		panic(err)
	}
	if note, ok := noteStore[key]; ok {
		noteUpdate.Id = note.Id
		noteUpdate.CreatedAt = note.CreatedAt
		delete(noteStore, key)
		noteStore[key] = noteUpdate
	} else {
		log.Printf("No encontramos el id %s", key)
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	if _, ok := noteStore[key]; ok {
		delete(noteStore, key)
	} else {
		log.Printf("No encontramos el id %s", key)
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	app := mux.NewRouter().StrictSlash(false)
	app.HandleFunc("/", getNotesHandler).Methods("GET")
	app.HandleFunc("/", postNoteHandler).Methods("POST")
	app.HandleFunc("/{id}", putNoteHandler).Methods("PUT")
	app.HandleFunc("/{id}", deleteNoteHandler).Methods("DELETE")
	server := &http.Server{
		Addr:           ":8080",
		Handler:        app,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("Listening")
	log.Println(server.ListenAndServe())
}
