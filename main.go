package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

var (
	stores map[string]string
	wg     sync.WaitGroup
	m      sync.Mutex
)

type Store struct {
	Id   string `json:"Id"`
	Item string `json:"Item"`
}

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

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/store", returnAllArticles)
	myRouter.HandleFunc("/store/{id}", Get)
	myRouter.HandleFunc("/store/{id}/{item}", Put)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	stores = map[string]string{"1": "soda", "2": "bread", "3": "bread"}
	wg.Add(1)
	go handleRequests()
	wg.Wait()
}
