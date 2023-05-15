package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/page/1.html")
	})
	r.GET("/haoel", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/articles/author/haoel/")
	})

	r.Static("/articles", "./content/articles")
	r.Static("/page", "./content/list")
	r.Static("/assets", "./assets")
	r.Static("/uploads", "./uploads")

	log.Println("[酷壳 Cool Shell Forever 电子存档]")
	log.Println()
	log.Println("芝兰生于深谷，不以无人而不芳")
	log.Println("君子修身养德，不以穷困而改志")
	log.Println()
	log.Println("感恩皓叔的无私分享，以及总在最需要的时刻愿意花时间给予指导和帮助。")
	log.Println()
	log.Println("工具使用：")
	log.Println("    如果你需要改变端口，可以使用环境变量 PORT 来指定端口，例如：")
	log.Println("    PORT=8080 ./forever-coolshell ")

	port := "8080"
	portEnv := os.Getenv("PORT")
	p, err := strconv.ParseInt(portEnv, 10, 64)
	if err != nil {
		log.Println("使用默认端口", port)
	} else {
		port = fmt.Sprintf("%d", p)
		log.Println("使用指定端口", port)
	}

	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
