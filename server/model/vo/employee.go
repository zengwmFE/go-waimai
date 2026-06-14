package vo

import "time"

type EmployeeLoginVO struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"userName"`
	Token    string `json:"token"`
}

type EmployeeUserVO struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
}

type EmployeeLoginResponseVO struct {
	Data EmployeeLoginVO `json:"data"`
	Code int             `json:"code"`
}

type EmployeeLoginError struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"msg"`
}

type EmployeeResponseVO struct {
	Data EmployeeUserVO `json:"data"`
	Code int            `json:"code"`
}

type EmployeeEditResponseVO struct {
	Data EmployRecord `json:"data"`
	Code int          `json:"code"`
}

type EmployRecord struct {
	ID         uint64    `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IDNumber   string    `json:"idNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser uint64    `json:"createUser"`
	UpdateUser uint64    `json:"updateUser"`
}

type EmployeeRecordsVO struct {
	Records []EmployRecord `json:"records"`
	Total   int64          `json:"total"`
}

type EmployeeListResponseVO struct {
	Code int               `json:"code"`
	Data EmployeeRecordsVO `json:"data"`
}
