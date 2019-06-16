package api

import (
	"time"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
)

// InitRouter set configs for gin engine router and set the APIs routes
func InitRouter(idb *InDB) *gin.Engine {
	router := gin.Default()
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		MaxAge:          50 * time.Second,
		Credentials:     true,
		ValidateHeaders: false,
	}))

	setRoutes(idb, router)

	return router
}

func setRoutes(idb *InDB, r *gin.Engine) {
	r.POST("/message/send", idb.Send)
	r.GET("/messages", idb.GetMessages)
	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	go msghandler()
}
