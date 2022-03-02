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
		Res.Code(http.StatusOK).Message("Up and Running").Data(gin.H{"app": "OK"}).Json(c)
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

	port, _ := Config.Int("App.Port")

	if port == 0 {
		port = 8080
	}

	Router.Run(fmt.Sprintf(":%d", port))
}
