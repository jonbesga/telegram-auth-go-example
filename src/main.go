package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	var host string = "0.0.0.0"
	var port int = 8080

	r := mux.NewRouter()
	private := r.PathPrefix("/").Subrouter()
	private.Use(loggingMiddleware)
	private.HandleFunc("/profile", ProfileHandler)

	public := r.PathPrefix("/").Subrouter()
	public.HandleFunc("/", IndexHandler)
	public.HandleFunc("/auth/telegram", AuthTelegramHandler)

	http.Handle("/", r)

	addr := fmt.Sprintf("%s:%d", host, port)
	fmt.Printf("Listening at http://%s\n", addr)

	server := &http.Server{Addr: addr}
	log.Fatal(server.ListenAndServe())
}
