package artifact

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Res ResponseBuilder

func LoadRoute() {

	gin.ForceConsoleColor()

	//gin.SetMode("debug")

	Router = gin.Default()

	//httpRouter.SetTrustedProxies([]string{"0.0.0.0"})

	Router.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running"})
	})
}

func LoadConfig() {
	Config = NewConfig()

	Config.Load()
}

func initializeLogger() LoggerBuilder {
	return NewLogger()
}

func ConnectDb() {
	Mongo = NewMongoDB()
}

func New() {
	LoadConfig()
	LoadRoute()
}

func Start() {
	ConnectDb()
	initializeLogger()
}

func Run() {
	defer Mongo.Client.Disconnect(Mongo.Ctx)

	port := Config.GetString("App.Port")

	if port == "" {
		port = "8080"
	}

	Router.Run(fmt.Sprintf(":%s", port))
}
