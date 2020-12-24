package api

type ReqGetUsers struct {
	Uid int64 `json:"uid" validate:"required"`
}

type RespGetUsers struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
