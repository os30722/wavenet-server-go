package vo

type CommentForm struct {
	ComentId int    `json:"comment_id"`
	PostId   int    `json:"post_id"`
	Content  string `json:"msg"`
	ParentId int    `json:"parent_id"`
}

type Comment struct {
	CommentId    int    `json:"comment_id"`
	Content      string `json:"msg"`
	Username     string `json:"username"`
	RepliesCount int    `json:"replies_count"`
}

type Like struct {
	Username string `json:"username"`
}
