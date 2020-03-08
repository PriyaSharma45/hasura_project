package router

import (
	"fmt"
	"hasura/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	GamerId     = "gamer_id"
	Name        = "name"
	SessionCode = "session_code"
)

func (sqlHandler *SQLHandler) CreateGameHandler(ctx *gin.Context) {

	gamerId := ctx.Query(GamerId)
	if gamerId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "gamer id wasn't provided",
		})
		return
	}
	name := ctx.Query(Name)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "name wasn't provided",
		})
		return
	}

	sessionid, err := service.CreateGameService(gamerId, name, sqlHandler.PostgresClient)
	if err != nil {
		fmt.Println("Err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't create game",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"session_id": sessionid,
		"message":    "game created",
	})
}

func (sqlHandler *SQLHandler) JoinGameHandler(ctx *gin.Context) {

	gamerId := ctx.Query(GamerId)
	if gamerId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "gamer id wasn't provided",
		})
		return
	}
	name := ctx.Query(Name)
	if name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "name wasn't provided",
		})
		return
	}

	sessionCode := ctx.Query(SessionCode)
	if sessionCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "session code wasn't provided",
		})
		return
	}

	err := service.JoinGameService(gamerId, name, sessionCode, sqlHandler.PostgresClient)
	if err != nil {
		fmt.Println("Err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't join game",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "joined game",
	})
}

func (sqlHandler *SQLHandler) StartGameHandler(ctx *gin.Context) {

	sessionCode := ctx.Query(SessionCode)
	if SessionCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "session code wasn't provided",
		})
		return
	}

	err := service.StartGameService(sessionCode, sqlHandler.PostgresClient)
	if err != nil {
		fmt.Println("Err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't start game",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "game started",
	})
}

func (sqlHandler *SQLHandler) GetWordHandler(ctx *gin.Context) {

	gamerId := ctx.Query(GamerId)
	if gamerId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "gamer id wasn't provided",
		})
		return
	}

	sessionCode := ctx.Query(SessionCode)
	if sessionCode == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "session code wasn't provided",
		})
		return
	}

	word, err := service.GetWordService(sessionCode, gamerId, sqlHandler.PostgresClient)
	if err != nil {
		fmt.Println("Err", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "couldn't join game",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "joined game",
		"word":    word,
	})
}
