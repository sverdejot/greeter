package bootstrap

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

var users map[int]string = map[int]string{
	1: "Samuel",
}

func Run() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /users/{id}", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		name, ok := users[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := struct{
			Name string `json:"name"`
		}{
			name,
		}

		json.NewEncoder(w).Encode(body)
	})

	server := http.Server{
		Addr: ":8081",

		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		IdleTimeout: 30 * time.Second,

		Handler: mux,
	}

	log.Println("starting app")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
