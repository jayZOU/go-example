package spider

type Article struct {
	Title string
	Url   string
	Desc  string `json:"desc,omitempty"`
}

type ArticleList map[string][]Article
