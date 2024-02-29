package vo

type Message struct {
	Msg string `json:"ms'g"`
}

type PageParams struct {
	Cursor   int
	PageSize int
}

type PageItem struct {
	TotalCounts int         `json:"total_count"`
	Items       interface{} `json:"items"`
}
