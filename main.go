package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

//go:embed all:assets
var Assets embed.FS

//go:embed all:content/haoel
var Author embed.FS

//go:embed all:content/articles
var Articles embed.FS

//go:embed all:content/list
var List embed.FS

//go:embed all:uploads
var Uploads embed.FS

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(gzip.Gzip(gzip.DefaultCompression))

	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/page/1.html")
	})

	assets, _ := fs.Sub(Assets, "assets")
	r.StaticFS("/assets", http.FS(assets))

	author, _ := fs.Sub(Author, "content/haoel")
	r.StaticFS("/haoel", http.FS(author))

	articles, _ := fs.Sub(Articles, "content/articles")
	r.StaticFS("/articles", http.FS(articles))

	page, _ := fs.Sub(List, "content/list")
	r.StaticFS("/page", http.FS(page))

	uploads, _ := fs.Sub(Uploads, "uploads")
	r.StaticFS("/uploads", http.FS(uploads))

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
