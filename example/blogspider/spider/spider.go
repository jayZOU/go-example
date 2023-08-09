package spider

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func Start() ArticleList {

	articleList := ArticleList{}

	articleList["张鑫旭"] = getZhangXinXu()
	articleList["阮一峰"] = getRuanYiFeng()
	articleList["Github Trending"] = getGithubTrending()
	return articleList
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

func getRuanYiFeng() []Article {
	var article []Article
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
		return article
	}

	//获取最新的文章和链接
	node := doc.Find(".entry-title").Find("a")
	articleUrl, _ := node.Attr("href")
	articleTitle := node.Text()
	article = append(article, Article{
		Title: articleTitle,
		Url:   articleUrl,
	})

	return article
}

func getZhangXinXu() []Article {
	layout := "2006年01月02日"
	var article []Article
	url := "https://www.zhangxinxu.com/wordpress/"
	doc := fetchHtml(url)

	doc.Find(".post").Each(func(i int, div *goquery.Selection) {
		url, exist := div.Find(".entry-title").Attr("href")
		title := div.Find(".entry-title").Text()
		date := div.Find(".date").Text()
		if exist {
			parse_time, _ := time.Parse(layout, date)
			timestamp := parse_time.Unix()
			if diffNow(timestamp) {

				article = append(article, Article{
					Title: title,
					Url:   url,
				})
			}
		}
	})
	return article
}

func getGithubTrending() []Article {
	var article []Article
	url := "https://github.com/trending"
	doc := fetchHtml(url)

	doc.Find(".Box-row").Each(func(i int, s *goquery.Selection) {
		url, _ := s.Find(".Box-row > .lh-condensed > .Link").Attr("href")
		url = fmt.Sprintf("https://github.com%s", url) //拼接访问URL

		title := s.Find(".Box-row > .lh-condensed > .Link").Text()
		title = strings.ReplaceAll(strings.ReplaceAll(title, " ", ""), "\n", "") //去除空格和换行

		desc := s.Find(".Box-row > p.color-fg-muted").Text()
		desc = strings.TrimSpace(strings.ReplaceAll(desc, "\n", ""))

		article = append(article, Article{
			Title: title,
			Url:   url,
			Desc:  desc,
		})
	})
	return article
}

func diffNow(timestamp int64) bool {
	now := time.Now()
	dayTime := 24 * 60 * 60

	return now.Unix()-timestamp < int64(dayTime)
}
