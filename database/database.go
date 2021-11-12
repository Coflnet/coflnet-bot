package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
)

type User struct {
	Id            string         `json:"id"`
	Role          string         `json:"role"`
	GoogleId      string         `json:"googleId"`
	MinecraftUuid sql.NullString `json:"minecraftUuid"`
	DiscordName   string         `json:"discordName"`
	PremiumExpires   string         `json:"premiumExpires"`
}

func UserById(id int) (User, error) {
	db := dbConnection()
	defer db.Close()

	var user User

	row := db.QueryRow("SELECT Id, GoogleId, MinecraftUuid, PremiumExpires from Users where Id = ?", id)

	if err := row.Scan(&user.Id, &user.GoogleId, &user.MinecraftUuid, &user.PremiumExpires); err != nil {
		log.Error().Err(err).Msgf("error getting user with id %d", id)
		return user, err
	}

	return user, nil
}

func AllUsers() ([]*User, error) {
	db := dbConnection()
	defer db.Close()

	users := []*User{}

	rows, err := db.Query("SELECT id, GoogleId, MinecraftUuid, PremiumExpires from Users order by Id")
	if err != nil {
		log.Error().Err(err).Msgf("there was an error when running SQL query")
		return nil, err
	}
	defer rows.Close()

	if err != nil {
		log.Error().Err(err).Msgf("error when fetching users all users database")
		return nil, err
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Id, &user.GoogleId, &user.MinecraftUuid, &user.PremiumExpires); err != nil {
			log.Error().Err(err).Msgf("error scanning through users")
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func dbConnection() *sql.DB {
	dsn := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@tcp(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
		os.Exit(1)
	}
	return db
}

func callMcConnect(user *User) error {
	url := fmt.Sprintf("http://%s/Connect/minecraft/%s", os.Getenv("MC_CONNECT"), user.MinecraftUuid)
	fmt.Println(url)
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(bytes, user); err != nil {
		return err
	}

	return nil
}
