package dto

type CategoryAddDTO struct {
	Name string `json:"name" binding:"required"`
	Type int    `json:"type" binding:"required"`
	Sort int    `json:"sort" json:"sort"`
}

type CategoryUpdateDTO struct {
	ID   uint64 `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Type int    `json:"type"`
	Sort int    `json:"sort" binding:"required"`
}

type CategoryPageQueryDTO struct {
	Name     string `form:"name" json:"name"`
	Type     *int   `form:"type" json:"type"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
}

type CategoryListQueryDTO struct {
	Type int `form:"type" json:"type"`
}
