package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	"github.com/gorilla/mux"
)

func Get(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	vars := mux.Vars(r)
	key := vars["id"]
	ok := true
	for k, v := range stores {
		if k == key {
			ok = false
			json.NewEncoder(w).Encode("Item of id `" + k + "` is `" + v + "`")
		}
	}
	if ok {
		json.NewEncoder(w).Encode("We don't have an item with id `" + key + "`")
	}

}

func Put(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	vars := mux.Vars(r)
	key := vars["id"]
	value := vars["item"]
	ok := true
	for k, v := range stores {
		if k == key {
			ok = false
			json.NewEncoder(w).Encode("We have updated item of ID `" + k + "` from `" + v + "` to `" + value + "`")
		}
	}
	if ok {
		json.NewEncoder(w).Encode("We have added item `" + value + "` with ID `" + key + "`")
	}
	stores[key] = value
	fmt.Println(stores)
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllArticles")
	json.NewEncoder(w).Encode(stores)
}
