package spider

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Start() map[string]Article {

	articleList := map[string]Article{}

	ruanyifeng, err := getRuanYiFeng()

	if err == nil {
		articleList["ruanyifeng"] = ruanyifeng
	}

	return articleList
}

func Test() map[string]struct{ a int } {
	return map[string]struct{ a int }{
		"test": {
			a: 1,
		},
	}
}

func fetchHtml(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch html url: %s, error: %s", url, err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatalf("load document error: %s", err)
	}

	return doc
}

func getRuanYiFeng() (Article, error) {
	var article = Article{
		Title: "",
		Url:   "",
	}
	url := "https://www.ruanyifeng.com/blog/"
	doc := fetchHtml(url)

	latest, exists := doc.Find(".published").Attr("title")
	if !exists {
		log.Fatal("获取最新文章发布时间失败")
	}

	parse_time, err := time.Parse(time.RFC3339, latest)
	if err != nil {
		log.Fatal("时间格式化失败")
	}

	//检查当天是否有更新文章
	now := time.Now()
	dayTime := 24 * 60 * 60
	if now.Unix()-parse_time.Unix() > int64(dayTime) {
		return article, errors.New("no latest articles")
	}

	//获取最新的文章和链接
	node := doc.Find(".entry-title").Find("a")
	articleUrl, _ := node.Attr("href")
	articleTitle := node.Text()
	article.Url = articleUrl
	article.Title = articleTitle

	return article, nil
}
