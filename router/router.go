package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type SQLHandler struct {
	PostgresClient *sql.DB
}

func GetRouterEngine(sqlHandler *SQLHandler) *gin.Engine {
	ginRouter := gin.Default()
	ginRouter.POST("/create", sqlHandler.CreateGameHandler)
	ginRouter.POST("/join", sqlHandler.JoinGameHandler)
	ginRouter.POST("/start", sqlHandler.StartGameHandler)
	ginRouter.POST("/get_word", sqlHandler.GetWordHandler)
	return ginRouter
}
