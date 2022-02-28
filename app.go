package artifact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Res ResponseBuilder

func loadRoute() {

	gin.ForceConsoleColor()

	//gin.SetMode("debug")

	Router = gin.Default()

	//httpRouter.SetTrustedProxies([]string{"0.0.0.0"})

	Router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running"})
	})
}

func loadConfig() {
	Config = NewConfig()

	Config.Load()
}

func initializeLogger() LoggerBuilder {
	return NewLogger()
}

func connectDb() {
	Mongo = NewMongoDB()
}

func init() {
	loadRoute()
	loadConfig()
	connectDb()
}

func Start() {
	initializeLogger()

}

func Run() {
	defer Mongo.Client.Disconnect(Mongo.Ctx)

	port := Config.GetString("App.Port")

	if port == "" {
		port = "8080"
	}

	Router.Run(fmt.Sprintf(":%d", port))
}
