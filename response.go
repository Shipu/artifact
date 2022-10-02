package artifact

import (
	"github.com/gin-gonic/gin"
	"reflect"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
	Raw     map[string]interface{} `json:"raw"`
}

type ResponseBuilder struct {
	CodeKeyName    string
	MessageKeyName string
	DataKeyName    string
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

func (response ResponseBuilder) Raw(raw map[string]interface{}) ResponseBuilder {
	response.Response.Raw = raw
	return response
}

func (response *ResponseBuilder) SetCodeKeyName(name string) *ResponseBuilder {
	response.CodeKeyName = name
	return response
}

func (response *ResponseBuilder) SetMessageKeyName(name string) *ResponseBuilder {
	response.MessageKeyName = name
	return response
}

func (response *ResponseBuilder) SetDataKeyName(name string) *ResponseBuilder {
	response.DataKeyName = name
	return response
}

func (response ResponseBuilder) Build() interface{} {
	if response.Response.Code == 0 {
		response.Code(200)
	}

	if response.DataKeyName == "" {
		response.SetDataKeyName("data")
	}

	if response.CodeKeyName == "" {
		response.SetCodeKeyName("status_code")
	}

	if response.MessageKeyName == "" {
		response.SetMessageKeyName("message")
	}

	if response.Response.Data != nil {
		data := reflect.TypeOf(response.Response.Data)
		switch data.Kind() {
		case reflect.Slice:
			if reflect.ValueOf(response.Response.Data).IsNil() {
				response.Response.Data = make([]interface{}, 0)
			}
		}
	}

	res := map[string]interface{}{response.CodeKeyName: response.Response.Code, response.MessageKeyName: response.Response.Message, response.DataKeyName: response.Response.Data}

	for key, value := range response.Response.Raw {
		res[key] = value
	}

	return res
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
	c.Header("Content-Type", "text/html")
	c.HTML(response.Response.Code, name, response.Build())
}

func (response ResponseBuilder) Xml(c *gin.Context) {
	c.Header("Content-Type", "text/xml")
	c.XML(response.Response.Code, response.Build())
}

func (response ResponseBuilder) Yaml(c *gin.Context) {
	c.Header("Content-Type", "text/yaml")
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
