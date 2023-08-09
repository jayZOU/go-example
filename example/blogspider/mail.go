package main

import (
	"log"

	"github.com/jayZOU/go-example/example/blogspider/email"
	htmlTemplate "github.com/jayZOU/go-example/example/blogspider/html"
	"github.com/jayZOU/go-example/example/blogspider/spider"
	"github.com/joho/godotenv"
)

func main() {
	//加载本地配置
	loadEnv()

	// 启动爬虫
	articleList := spider.Start()

	log.Print(articleList)

	if articleLen(articleList) == 0 {
		log.Print("article len: 0")
		return
	}

	html := htmlTemplate.GenerateHtml(articleList)
	email.Send(html)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func articleLen(articleList spider.ArticleList) int {
	articleLen := 0
	for _, articles := range articleList {
		articleLen += len(articles)
	}
	return articleLen
}
