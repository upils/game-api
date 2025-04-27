package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/upils/game-api/model"
)

func (s *gameServer) GetGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var game = model.Game{ID: uint(id)}

		result := s.db.WithContext(r.Context()).Model(&model.Game{}).Preload("Platforms").First(&game)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jGame, err := json.Marshal(&game)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jGame)
	})
}
func (s *gameServer) ListGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var games = []model.Game{}

		result := s.db.WithContext(r.Context()).Model(&model.Game{}).Preload("Platforms").Find(&games)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jGames, err := json.Marshal(&games)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jGames)
	})
}

func (s *gameServer) CreateGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var game model.Game
		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(game.Name) == 0 {
			http.Error(w, "name cannot be empty", http.StatusBadRequest)
			return
		}

		if game.Ratings > 20 {
			http.Error(w, "ratings cannot be greated than 20", http.StatusBadRequest)
			return
		}

		result := s.db.WithContext(r.Context()).Create(&game)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jGame, err := json.Marshal(&game)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jGame)
	})
}

func (s *gameServer) UpdateGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var game model.Game
		err := json.NewDecoder(r.Body).Decode(&game)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result := s.db.WithContext(r.Context()).First(&game)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		result = s.db.WithContext(r.Context()).Save(&game)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}

		jGame, err := json.Marshal(&game)
		if result.Error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(jGame)
	})
}

func (s *gameServer) DeleteGameHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := s.db.WithContext(r.Context()).Delete(&model.Game{}, id)
		if result.Error != nil {
			http.Error(w, result.Error.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
}
