package main

import (
	"net/http"
)

func main() {
	api := &api{addr: ":8080"}

	// servemux başlat / mux bir routerdır
	mux := http.NewServeMux() // router

	srv := &http.Server{ // handler interfaceimiz
		Addr:    api.addr,
		Handler: mux,
	}

	mux.HandleFunc("GET /users", api.getUsersHandler)
	mux.HandleFunc("POST /users", api.createUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
