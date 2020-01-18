package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterHighscoresEndpoints(router *mux.Router) {
	router.HandleFunc("/highscores", HighscoresHandler)
}

func HighscoresHandler(writer http.ResponseWriter, request *http.Request) {
	var mode = request.FormValue("mode")

	switch mode {

	case "orbpoints":
		HighscoresOrbPointsHandler(writer)
		return

	default:
		writer.WriteHeader(http.StatusBadRequest)
	}
}

func HighscoresOrbPointsHandler(writer http.ResponseWriter) {
	var entries []struct {
		tableName   struct{} `pg:"account"`
		DisplayName string   `pg:"display_name" json:"displayName"`
		OrbPoints   int      `pg:"orb_points" json:"orbPoints"`
	}
	err := postgres.Model(&entries).
		Order("orb_points DESC").
		Limit(10).
		Select()

	if err != nil {
		DefaultQueryErrorHandler(err, writer)
	} else {
		json, _ := json.Marshal(entries)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	}
}
