package vo

import "sky-take-out-go/model/entity"

type CategoryVO struct{}

type CategoryResError struct {
	Code   uint64 `json:"code"`
	ErrMsg string `json:"errmsg"`
}

type CatorySuccessData struct {
	ID uint64
}

type CategoryResSuccess struct {
	Code uint64            `json:"code" binding:"required"`
	Data CatorySuccessData `json:"data" binding:"required"`
}

type CategoryListSuccessData struct {
	Records []entity.Category `json:"records" binding:"required"`
	Total   int64             `json:"total" binding:"required"`
}

type CategoryResListSuccess struct {
	Code uint64                  `json:"code" binding:"required"`
	Data CategoryListSuccessData `json:"data" binding:"required"`
}

type CategoryListResData struct {
	Code uint64             `json:"code" binding:"required"`
	Data *[]entity.Category `json:"data" binding:"required"`
}
