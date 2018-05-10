package main

import (
	"net/http"

	"github.com/bjorm/wasgeit"
	log "github.com/sirupsen/logrus"
)

func main() {
	store := &wasgeit.Store{}
	dbErr := store.Connect()

	if dbErr != nil {
		panic(dbErr)
	}
	defer store.Close()

	server := wasgeit.NewServer(store)
	http.HandleFunc("/agenda", server.ServeAgenda)
	http.HandleFunc("/news", server.ServeNews)

	log.Info("Serving..")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
