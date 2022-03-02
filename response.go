package artifact

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseBuilder struct {
	Response
}

func (response ResponseBuilder) Code(status int) ResponseBuilder {
	response.Response.Code = status
	return response
}

func (response ResponseBuilder) Message(message string) ResponseBuilder {
	response.Response.Message = message
	return response
}

func (response ResponseBuilder) Data(data interface{}) ResponseBuilder {
	response.Response.Data = data
	return response
}

func (response ResponseBuilder) Build() interface{} {
	if response.Response.Code == 0 {
		response.Code(200)
	}

	return map[string]interface{}{"code": response.Response.Code, "message": response.Response.Message, "data": response.Response.Data}
}

func (response ResponseBuilder) Json(c *gin.Context) {
	c.JSON(response.Response.Code, response.Build())
}

func (response ResponseBuilder) PureJSON(c *gin.Context) {
	c.PureJSON(response.Response.Code, response.Build())
}

func (response ResponseBuilder) JsonP(c *gin.Context) {
	c.JSONP(response.Response.Code, response.Build())
}

func (response ResponseBuilder) AsciiJSON(c *gin.Context) {
	c.AsciiJSON(response.Response.Code, response.Build())
}

func (response ResponseBuilder) IndentedJSON(c *gin.Context) {
	c.IndentedJSON(response.Response.Code, response.Build())
}

func (response ResponseBuilder) Html(c *gin.Context, name string) {
	c.HTML(response.Response.Code, name, response.Build())
}

func (response ResponseBuilder) Xml(c *gin.Context) {
	c.XML(response.Response.Code, response.Build())
}

func (response ResponseBuilder) Yaml(c *gin.Context) {
	c.YAML(response.Response.Code, response.Build())
}

func (response ResponseBuilder) ProtoBuf(c *gin.Context) {
	c.ProtoBuf(response.Response.Code, response.Build())
}

func (response ResponseBuilder) AbortWithStatusJSON(c *gin.Context) {
	c.AbortWithStatusJSON(response.Response.Code, response.Build())
}

func (response ResponseBuilder) Redirect(c *gin.Context, location string) {
	c.AbortWithStatusJSON(response.Response.Code, location)
}

func (response ResponseBuilder) Abort(c *gin.Context) {
	c.Abort()
}

func (response ResponseBuilder) AbortWithError(c *gin.Context, err error) {
	c.AbortWithError(response.Response.Code, err)
}
