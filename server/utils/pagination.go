package utils

type PageQuery struct {
	Page     int `form:"page" json:"page"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

func (q *PageQuery) Offset() int {
	if q.Page <= 0 {
		q.Page = 1
	}
	if q.PageSize <= 0 {
		q.PageSize = 10
	}
	return (q.Page - 1) * q.PageSize
}

func (q *PageQuery) Limit() int {
	if q.PageSize <= 0 {
		return 10
	}
	return q.PageSize
}

type PageResult struct {
	Total   int64       `json:"total"`
	Records interface{} `json:"records"`
}
