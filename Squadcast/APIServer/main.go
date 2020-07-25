package main

import (
	"Squadcast/APIServer/handler"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/incident/create", handler.HandledCreateRequest).Methods("POST")
	router.HandleFunc("/api/incident/resolve", handler.HandledResolveRequest).Methods("POST")
	router.HandleFunc("/api/incident/ack", handler.HandledAckRequest).Methods("POST")
	router.HandleFunc("/api/incident/comment", handler.HandledCommentRequest).Methods("POST")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Print("Not able to connect on the port 5000", err)
		return
	}

}
