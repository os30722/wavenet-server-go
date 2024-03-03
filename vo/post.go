package vo

type Post struct {
	PostId    int    `json:"post_id"`
	Title     string `json:"title"`
	Url       string `json:"url"`
	UserName  string `json:"username"`
	Likes     int    `json:"likes"`
	UserLiked bool   `json:"user_liked"`
}

type PostUpload struct {
	Title       string
	Description string
	FileName    string
	UserId      int
	Extention   string
}
