package vo

type Post struct {
	PostId int    `json:"post_id"`
	Title  string `json:"title"`
	Url    string `json:"url"`
	User   string `json:"user"`
}

type PostUpload struct {
	Title       string
	Description string
	FileName    string
}
