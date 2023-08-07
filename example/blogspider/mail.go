package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"

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

	if len(articleList) == 0 {
		log.Print("article len: 0")
		return
	}

	html := generateHtml(htmlTemplate.HTML, articleList)

	email.Send(html)

}

func generateHtml(html string, articleList map[string]spider.Article) string {

	for auther, article := range articleList {
		field := reflect.TypeOf(article)
		valueOf := reflect.ValueOf(article)
		fieldLen := field.NumField()

		for i := 0; i < fieldLen; i++ {
			key := field.Field(i).Name
			value := valueOf.Field(i)

			mark := fmt.Sprintf("{{%s.%s}}", auther, key)
			html = strings.ReplaceAll(html, mark, value.String())
		}

	}
	return html
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
