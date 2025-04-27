package main

import (
	"log"
	"net/http"
	"time"

	"github.com/upils/game-api/api"
	"github.com/upils/game-api/db"
)

func main() {
	log.Println("Open connection to DB...")
	db := db.NewDB()

	gameAPI := api.NewGameServer(db)

	mux := http.NewServeMux()
	mux.Handle("GET /game/{id}", logger(gameAPI.GetGameHandler()))
	mux.Handle("GET /games/", logger(gameAPI.ListGameHandler()))
	mux.Handle("POST /game/", logger(gameAPI.CreateGameHandler()))
	mux.Handle("PUT /game/", logger(gameAPI.UpdateGameHandler()))
	mux.Handle("DELETE /game/{id}", logger(gameAPI.DeleteGameHandler()))

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Server listening on port 8080...")
	log.Fatal(srv.ListenAndServe())
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request URI: %s, Method: %s", r.RequestURI, r.Method)
		next.ServeHTTP(w, r)
	})
}
