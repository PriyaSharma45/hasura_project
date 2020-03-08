package service

import (
	"database/sql"
	"errors"
	"fmt"
	"hasura/models"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func CreateGameService(gamerId string, name string, postGres *sql.DB) (string, error) {

	session_code := StringWithCharset(8, charset)

	_, err := postGres.Exec(`INSERT INTO Game(gamer_id, session_code, name, is_primary_player)
	VALUES($1, $2, $3, $4)`, gamerId, session_code, name, true)

	if err != nil {
		return "", err
	}
	return session_code, nil
}

func JoinGameService(gamerId string, name string, sessionCode string, postGres *sql.DB) error {

	var session string
	err := postGres.QueryRow(`SELECT session_code FROM Game WHERE session_code=$1`, sessionCode).Scan(&session)
	if err != nil {
		fmt.Println("error while getting session code ", err)
		return err
	}

	if session != sessionCode {
		errors.New("create game to join")
	}

	_, err = postGres.Exec(`INSERT INTO Game(gamer_id, session_code, name, is_primary_player)
	VALUES($1, $2, $3, $4)`, gamerId, sessionCode, name, false)

	if err != nil {
		return err
	}
	return nil
}

func StartGameService(sessionCode string, postGres *sql.DB) error {
	var gamerIds []string
	var words []string
	var session string

	err := postGres.QueryRow(`SELECT session_code FROM Game WHERE session_code=$1`, sessionCode).Scan(&session)
	if err != nil {
		fmt.Println("error while getting session code ", err)
		return err
	}

	if session != sessionCode {
		return errors.New("session doesn't exist")
	}

	rows, err := postGres.Query(`SELECT gamer_id FROM Game WHERE session_code=$1`, sessionCode)
	if err != nil {
		fmt.Println("error while getting spy ", err)
		return err
	}

	for rows.Next() {
		var gamerId string

		err = rows.Scan(&gamerId)
		if err != nil {
			return err
		}
		gamerIds = append(gamerIds, gamerId)

	}

	if len(gamerIds) < 3 {
		return errors.New("minimum three people are required to start the game")
	}

	spyId := gamerIds[rand.Intn(len(gamerIds))]

	wordRows, err := postGres.Query(`SELECT word FROM words`)
	if err != nil {
		fmt.Println("error while getting word ", err)
		return err
	}

	for wordRows.Next() {
		var word string

		err = wordRows.Scan(&word)
		if err != nil {
			return err
		}
		words = append(words, word)
	}

	word := words[rand.Intn(len(words))]

	_, err = postGres.Exec(`INSERT INTO Session(session_id, is_started, word, spy_gamer_id)
	VALUES($1, $2, $3, $4)`, sessionCode, true, word, spyId)
	if err != nil {
		return err
	}

	return nil
}

func GetWordService(sessionCode string, gamerId string, postGres *sql.DB) (string, error) {

	var sessionStruct models.Session
	err := postGres.QueryRow(`SELECT * FROM Session WHERE session_id=$1`, sessionCode).
		Scan(&sessionStruct.SessionCode, &sessionStruct.IsStarted, &sessionStruct.Word, &sessionStruct.SpyGamerId)
	if err != nil {
		fmt.Println("error while getting word ", err)
		return "", err
	}

	if !sessionStruct.IsStarted {
		return "", errors.New("game hasn't started")
	}

	if sessionStruct.SpyGamerId == gamerId {
		return "", nil
	}

	fmt.Println("sessionStruct", sessionStruct)
	return sessionStruct.Word, nil
}
