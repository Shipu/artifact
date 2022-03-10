# Artifact

Artifact is a web framework written in Go (Golang) based on [Gin](https://github.com/gin-gonic/gin), [MongoDB](https://github.com/mongodb/mongo-go-driver).

Example Repo: [Golang Gin Boilerplate](https://github.com/Shipu/golang-gin-boilerplate)

## Installation
```go
go get -u github.com/shipu/artifact
```

## Quick Start
```shell
$ cat .env
```
```dotenv
APP_NAME=app
APP_ENV=local
APP_DEBUG=true
APP_URL=http://localhost
APP_PORT=8080
```
```shell
# assume the following codes in example.go file
$ cat example.go
```

```go
package main

import (
	. "github.com/shipu/artifact"
	"github.com/gin-gonic/gin"
)

type AppConfig struct {
	Name        string `mapstructure:"APP_NAME" default:"Artifact"`
	Environment string `mapstructure:"APP_ENV" default:"local"`
	Debug       bool   `mapstructure:"APP_DEBUG" default:"true"`
	Url         string `mapstructure:"APP_URL"  default:"http://localhost"`
	Port        int    `mapstructure:"APP_PORT" default:"8098"`
	TimeZone    string `mapstructure:"APP_TIMEZONE"  default:"UTC"`
	Locale      string `mapstructure:"APP_LOCALE"  default:"en"`
	GinMode     string `mapstructure:"GIN_MODE" default:"debug"`
}


func main() {
	// Initialize the application
    New()

    Config.AddConfig("App", new(AppConfig)).Load()
    
    // artifact.Start() // Database connection will be established here
    
    Router.GET("/", func(c *gin.Context) {
        data := map[string]interface{}{
            "app": Config.GetString("App.Name"),
        }
    
        //or
        //data := gin.H{
        //	"message": "Hello World",
        //}
    
        Res.Status(200).
            Message("success").
            Data(data).Json(c)
    })
    
    Run()
}
```

```shell
# run example.go and visit 0.0.0.0:8080 (for windows "localhost:8080") on browser
$ go run example.go
```

## Crud Generator

```go
go run ./art crud package_name crud_module_name
```

for example:
```shell
$ cat art/main.go
```

```go
package main

import "github.com/shipu/artifact/cmd"

func main() {
    cmd.Execute()
}
```
Run below command to generate crud.
```go 
go run ./art crud github.com/shipu/golang-gin-boilerplate notice
``` 

Below Folder structure will generate:
```go
src/notice
├── controllers
│   └── notice_controller.go
├── models
│   └── notice.go
├── routes
│   └── api.go
└── services
└── notice_service.go
```

More information about crud generator can be found in [Golang Gin Boilerplate](https://github.com/Shipu/golang-gin-boilerplate)

## Config :

Suppose your config is `config/db.go`:
```go
package config

type DatabaseConfig struct {
    Username   string `mapstructure:"DB_USER" default:""`
    Password   string `mapstructure:"DB_PASS" default:""`
    Host       string `mapstructure:"DB_HOST" default:""`
    Port       string `mapstructure:"DB_PORT" default:""`
    Database   string `mapstructure:"DB_DATABASE" default:""`
    Connection string `mapstructure:"DB_CONNECTION" default:""`
}
```
and your `.env` is:
```dotenv
DB_CONNECTION=mongodb
DB_HOST=mongodb.host
DB_PORT=
DB_USER=user
DB_PASS=password
DB_DATABASE=collection
```

For initialization `DatabaseConfig` config.

```go
artifact.Config.AddConfig("DB", new(DatabaseConfig)).Load()
```

### To get config:
```go
artifact.Config.GetString("DB.Host")
```

Config Method List:
```go
GetString("key")
GetInt("key")
Get("key")
```

## Route

```go
artifact.Router.GET("/", func(c *gin.Context) {
    data := map[string]interface{}{
        "app": Config.GetString("App.Name"),
    }
    
    //or
    //data := gin.H{
    //	"message": "Hello World",
    //}
    
    Res.Status(200).
        Message("success").
        Data(data).Json(c)
})
```

```go
artifact.Router.GET("/someGet", getting)
artifact.Router.POST("/somePost", posting)
artifact.Router.PUT("/somePut", putting)
artifact.Router.DELETE("/someDelete", deleting)
artifact.Router.PATCH("/somePatch", patching)
artifact.Router.HEAD("/someHead", head)
artifact.Router.OPTIONS("/someOptions", options)
```

And all [Gin router](https://github.com/gin-gonic/gin/edit/master/README.md#using-get-post-put-patch-delete-and-options) support.

## Response
In [Gin](https://github.com/gin-gonic/gin)

Where `c` is the `*gin.Context` context.

```go
data := map[string]interface{}{
    "app": "Golang",
}
c.JSON(200, gin.H{
    "status_code":  200,
    "message": "Success",
    "data": data,
})
```
In artifact
```go
data := map[string]interface{}{
    "app": "Golang",
}

Res.Code(200).
    Message("Success").
    Data(data).
    Json(c)
```

for set custom key value in response
```go
paginate := your paginate data

Res.Code(200).
    Message("Success").
    Data(data).
    Raw(map[string]interface{}{
        "meta": paginate,
    }).
    Json(c)
```

`Res` Api Methods:
```go
Json
PureJSON
JsonP
AsciiJSON
IndentedJSON
Html
Xml
Yaml
ProtoBuf
AbortWithStatusJSON
Abort
AbortWithError
Redirect
```

## Mongo Collection

```go
var TodoCollection artifact.MongoCollection = artifact.Mongo.Collection("todos")

TodoCollection.Find(bson.M{})
```

All [Go Mongo Driver](https://docs.mongodb.com/drivers/go/current/) Support.