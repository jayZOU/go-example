package html

import (
	"fmt"

	"github.com/jayZOU/go-example/example/blogspider/spider"
)

func GenerateHtml(articleList spider.ArticleList) string {
	var html string
	for auther, articles := range articleList {
		if len(articles) == 0 {
			continue
		}
		articleListTemp := ``
		for _, article := range articles {
			articleListTemp += fmt.Sprintf("<p><a href='%s'>%s</a><br><span>%s</span><br></p>", article.Url, article.Title, article.Desc)
		}
		html += fmt.Sprintf(`
			<h2>%s</h2>
			<div>
				%s
			</div>
		`, auther, articleListTemp)
	}
	return fmt.Sprintf(`
        <!DOCTYPE html>
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta http-equiv="X-UA-Compatible" content="IE=edge">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title></title>
        </head>
        <body>
            <div>
                %s
            </div>
        </body>
        </html>
    `, html)
}
