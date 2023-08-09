package spider

type Article struct {
	Title string
	Url   string
}

type ArticleList map[string][]Article
