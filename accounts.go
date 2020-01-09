package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterAccountEndpoints(router *mux.Router) {
	router.HandleFunc("/accounts", AccountsHandler)
	router.HandleFunc("/accounts/{uuid}", AccountDetailHandler)
	router.HandleFunc("/accounts/{uuid}/achievements", AccountAchievementsHandler)
}

func AccountsHandler(writer http.ResponseWriter, request *http.Request) {
	displayName := request.FormValue("displayName")

	if displayName == "" {
		writer.WriteHeader(http.StatusBadRequest)
	} else {
		var account Account

		err := postgres.Model(&account).
			Where("display_name = ?", displayName).
			Select()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			log.Print("Error while querying database => ", err)
		} else if account.DisplayName != "" {
			json, _ := json.Marshal(account)
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(json)
		} else {
			writer.WriteHeader(http.StatusNotFound)
		}
	}
}

func AccountDetailHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	uuid := vars["uuid"]

	var account Account

	err := postgres.Model(&account).
		Where("uuid = ?", uuid).
		Select()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print("Error while querying database => ", err)
	} else if account.DisplayName != "" {
		json, _ := json.Marshal(account)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func AccountAchievementsHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	uuid := vars["uuid"]
	gameId := request.FormValue("gameId")
	var achievements []PlayerGameAchievement

	query := postgres.Model(&achievements).
		Join("INNER JOIN account acc").
		JoinOn("account_id = acc.id").
		Where("acc.uuid = ?", uuid)

	if gameId != "" {
		query.Where("game_id = ?", gameId)
	}
	err := query.Select()

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print("Error while querying database => ", err)
	} else if achievements != nil {
		json, _ := json.Marshal(achievements)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
