package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func RegisterAccountEndpoints(router *mux.Router) {
	router.HandleFunc("/accounts", AccountsHandler).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{uuid}", AccountDetailHandler).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{uuid}/achievements", AccountAchievementsHandler).Methods(http.MethodGet)
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
			DefaultQueryErrorHandler(err, writer)
		} else {
			json, _ := json.Marshal(account)
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(json)
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
		DefaultQueryErrorHandler(err, writer)
	} else {
		json, _ := json.Marshal(account)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
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
		DefaultQueryErrorHandler(err, writer)
	} else {
		json, _ := json.Marshal(achievements)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(json)
	}
}
