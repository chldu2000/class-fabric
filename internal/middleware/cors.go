package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// CORSMiddleware 跨域请求中间件
func CORSMiddleware() gin.HandlerFunc {
	//cfg := config.C.CORS
	//fmt.Println("到跨域了")
	//return cors.New(cors.Config{
	//	AllowOrigins:     cfg.AllowOrigins,
	//	AllowMethods:     cfg.AllowMethods,
	//	AllowHeaders:     cfg.AllowHeaders,
	//	AllowCredentials: cfg.AllowCredentials,
	//	MaxAge:           time.Second * time.Duration(cfg.MaxAge),
	//})
	return func(context *gin.Context) {
		fmt.Println("到跨域了")
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}
		if origin != "" {
			fmt.Println("到origin这里了")
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}
		if method == "OPTIONS" {
			fmt.Println("处理option")
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		fmt.Println("跨域处理结束")
		context.Next()
	}

	//return func(c *gin.Context) {
	//	method := c.Request.Method
	//	origin := c.Request.Header.Get("Origin") //请求头部
	//	if origin != ""{
	//		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
	//		// 接受指定域的请求，可以使用*不加以限制，但不安全
	//		c.Header("Access-Control-Allow-Origin", "*")
	//		//请求的允许的头字段与权限控制
	//		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	//		//请求类型
	//		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,PUT")
	//		const value = "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type"
	//		c.Header("Access-Control-Expose-Headers", value)
	//		//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	//		c.Header("Access-Control-Allow-Credentials", "true")
	//	}
	//
	//	// 放行所有OPTIONS方法
	//	if method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//	}
	//	c.Next()
	//}

	//return func(c *gin.Context) {
	//	method := c.Request.Method
	//	// 接受指定域的请求，可以使用*不加以限制，但不安全
	//	c.Header("Access-Control-Allow-Origin", "127.0.0.1")
	//	//请求的允许的头字段与权限控制
	//	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	//	//请求类型
	//	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	//	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	//	//是否允许后续请求携带认证信息,该值只能是true,否则不返回
	//	c.Header("Access-Control-Allow-Credentials", "true")
	//	// 放行所有OPTIONS方法
	//	if method == "OPTIONS" {
	//		c.AbortWithStatus(http.StatusNoContent)
	//	}
	//	c.Next()
	//}

}
