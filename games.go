package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func RegisterGamesEndpoints(router *mux.Router) {
	router.HandleFunc("/games", GamesHandler)
	router.HandleFunc("/games/{id}/achievements", GameAchievementsHandler)
}

func GamesHandler(writer http.ResponseWriter, request *http.Request) {
	var games []Game

	err := postgres.Model(&games).
		Select()

	if err != nil {
		DefaultQueryErrorHandler(err, writer)
	} else {
		writer.Header().Set("Content-Type", "application/json")
		json, _ := json.Marshal(games)
		writer.Write(json)
	}
}

func GameAchievementsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	gameId, err := strconv.Atoi(vars["id"])

	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	var achievements []GameAchievement

	err = postgres.Model(&achievements).
		Where("game_id = ?", gameId).
		Select()

	if err != nil {
		DefaultQueryErrorHandler(err, writer)
	} else {
		json, _ := json.Marshal(achievements)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	}
}
