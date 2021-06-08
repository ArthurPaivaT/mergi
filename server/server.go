package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/ArthurPaivaT/mergi/handlers"
	"github.com/gorilla/mux"
)

func Start(resCh chan error) {
	fmt.Println("Starting Server...")

	router := mux.NewRouter()

	router.Handle("/applyBW", (http.HandlerFunc(handlers.ApplyBW))).Methods("POST")

	router.Use(mux.CORSMethodMiddleware(router))
	// needs to be non-empty

	fmt.Println("Listening on Port :1212")
	err := http.ListenAndServe(":1212", router)
	if err != nil {
		fmt.Println("Could not start server:", err)
	}

	resCh <- errors.New("Error ao iniciar servidor")

}
