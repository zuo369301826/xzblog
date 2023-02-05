package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Run() {

	router := gin.Default()

	// http://localhost:8099/
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, Geektutu")
	})

	// http://localhost:8099/ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// http://localhost:8099/user/xz
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// http://localhost:8099/welcome?firstname=small&lastname=left
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.DefaultQuery("lastname", "Guest")
		lastname2 := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s %s", firstname, lastname, lastname2)
	})

	// http://localhost:8099/form_post
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	// http://localhost:8099/index/template
	router.LoadHTMLGlob("data/*")
	//router.LoadHTMLFiles("data/template.html")
	router.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.Run(":8099")
}
