package dto

type EmployeeLoginDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeePageQueryDTO struct {
	Name     string `form:"name" json:"name"`
	Page     int    `form:"page" json:"page"`
	PageSize int    `form:"pageSize" json:"pageSize"`
}

type EmployeeAddDTO struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	IDNumber string `json:"idNumber" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type EmployeeUpdateDTO struct {
	ID       uint64 `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
	IDNumber string `json:"idNumber" binding:"required"`
}

type EmployeeQueryById struct {
	ID uint64 `uri:"id" binding:"required"`
}
type EmployeeQueryByStatus struct {
	Status int    `uri:"status" binding:"required"`
	ID     uint64 `form:"id" binding:"required"`
}

type EmployeeEditPasswordDTO struct {
	OldPassword string `json:"oldPassword" binding:"required"`
	NewPassword string `json:"newPassword" binding:"required"`
}

type EmployeeStatusDTO struct {
	Status int `json:"status" binding:"required"`
}
