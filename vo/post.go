package vo

type Post struct {
	PostId int    `json:"post_id"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	User   string `json:"user"`
}
