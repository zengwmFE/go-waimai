package vo

type UserLoginVO struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Token  string `json:"token"`
}
