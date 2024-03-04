package vo

type Comment struct {
	ComentId int    `json:"comment_id"`
	PostId   int    `json:"post_id"`
	Content  string `json:"msg"`
	ParentId int    `json:"parent_id"`
}
