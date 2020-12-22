package protocol

type ReqGetUsers struct {
	Uid uint64 `json:"uid" validate:"required"`
}

type RespGetUsers struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}
