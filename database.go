package main

import (
	"github.com/go-pg/pg/v9"
	"log"
	"net/http"
	"os"
	"time"
)

var postgres *pg.DB

func init() {
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	username := os.Getenv("DATABASE_USERNAME")
	password := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")
	postgres = pg.Connect(&pg.Options{
		Addr:     host + ":" + port,
		User:     username,
		Password: password,
		Database: database,
	})
}

func DefaultQueryErrorHandler(err error, writer http.ResponseWriter) {

	if err == pg.ErrNoRows {
		writer.WriteHeader(http.StatusNotFound)
	} else {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Print("Error while querying database => ", err)
	}
}

type Account struct {
	tableName   struct{} `pg:"account"`
	Uuid        string   `json:"uuid" pg:"uuid"`
	DisplayName string   `json:"displayName" pg:"display_name"`
	OrbCoins    int      `json:"orbCoins" pg:"orb_coins"`
	OrbPoints   int      `json:"orbPoints" pg:"orb_points"`
}

type Game struct {
	tableName              struct{} `pg:"game"`
	Id                     int      `json:"id" pg:"id"`
	JagexName              string   `json:"jagexName" pg:"jagex_name"`
	FancyName              string   `json:"fancyName" pg:"fancy_name"`
	LoginEnabled           bool     `json:"loginEnabled" pg:"login_enabled"`
	ObtainableAchievements int      `json:"obtainableAchievements" pg:"obtainable_achievements"`
	ObtainableOrbCoins     int      `json:"obtainableOrbCoins" pg:"obtainable_orb_coins"`
	ObtainableOrbPoints    int      `json:"obtainableOrbPoints" pg:"obtainable_orb_points"`
}

type GameAchievement struct {
	tableName     struct{} `pg:"game_achievement"`
	GameId        int      `json:"gameId" pg:"game_id"`
	AchievementId int      `json:"achievementId" pg:"achievement_id"`
	Name          string   `json:"name" pg:"name"`
	Criteria      string   `json:"criteria" pg:"criteria"`
	OrbCoins      int      `json:"orbCoins" pg:"orb_coins"`
	OrbPoints     int      `json:"orbPoints" pg:"orb_points"`
	Obtainable    bool     `json:"obtainable" pg:"obtainable"`
}

type PlayerGameAchievement struct {
	tableName       struct{}  `pg:"player_game_achievement"`
	Id              int       `pg:"achievement_id" json:"id"`
	GameId          int       `pg:"game_id" json:"gameId"`
	UnlockTimestamp time.Time `pg:"unlock_timestamp" json:"unlockTimestamp"`
}
