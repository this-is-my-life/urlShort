package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"urlShort/api"
)

func main() {
	PORT := ":8080"
	if value, has := os.LookupEnv("urlShort"); has {
		PORT = ":" + value
	}

	mux := &http.ServeMux{}
	srv := &http.Server{
		Addr:         PORT,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	api.ConnectDB()

	src := http.FileServer(http.Dir("./src"))
	src = http.StripPrefix("/src/", src)

	mux.Handle("/src/", src)
	mux.HandleFunc("/", api.R)

	fmt.Println("urlShort is now on http://localhost" + PORT)

	err := srv.ListenAndServe()
	erring(err)
}

func erring(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
