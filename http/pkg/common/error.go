package common

import "fmt"

type Err struct {
	ErrorNo  int
	ErrorMsg string
}

func (this Err) Error() string {
	return this.ErrorMsg
}

func (this Err) Errorf(errmsg ...string) Err {
	this.ErrorMsg = fmt.Sprintf(this.ErrorMsg, errmsg)
	return this
}

//业务错误码
var (
	ERR_SUC             = Err{ErrorNo: 0, ErrorMsg: "OK"}
	ERR_UNKNOWN         = Err{ErrorNo: -1, ErrorMsg: "未知错误"}
	ERR_PARAM           = Err{ErrorNo: 1000, ErrorMsg: "解析参数错误:%s"}
	ERR_VALIDATOR_PARAM = Err{ErrorNo: 1001, ErrorMsg: "请求参数错误:%s"}
)

//数据库错误码
var (
	ERR_DAO_NOT_FOUND = Err{ErrorNo: 0, ErrorMsg: "查询不到数据"}
)
