package http

import (
	"github.com/gin-gonic/gin"
	"go-framework/http/pkg/common"
)

//获取 json 格式的请求数据
func GetBodyParam(c *gin.Context, keyStruct interface{}) (err error) {
	if err = c.ShouldBind(keyStruct); err != nil {
		return common.ERR_PARAM.Errorf(err.Error())
	}
	if err = common.Validator(keyStruct); err != nil {
		return common.ERR_VALIDATOR_PARAM.Errorf(err.Error())
	}
	return
}

// ResponseSuccess 成功返回AccountMessageContentDataStruct
func ResponseSuccess(c *gin.Context, data interface{}) {
	ret := map[string]interface{}{
		"statusCode": common.ERR_SUC.ErrorNo,
		"msg":        common.ERR_SUC.ErrorMsg,
		"data":       data,
	}

	RenderJson(c, ret)
	return
}

// 返回data数据及错误码
func ResponseErrorCodeAndData(c *gin.Context, err *common.Err, data interface{}) {
	ret := map[string]interface{}{
		"statusCode": err.ErrorNo,
		"msg":        err.ErrorMsg,
		"data":       data,
	}

	RenderJson(c, ret)
	return
}

// 失败返回
func ResponseError(c *gin.Context, err error) {
	commonErr, ok := err.(common.Err)
	if !ok {
		commonErr = common.ERR_UNKNOWN
	}

	ret := map[string]interface{}{
		"statusCode": commonErr.ErrorNo,
		"msg":        commonErr.ErrorMsg,
		"data":       []interface{}{},
	}

	RenderJson(c, ret)
	return
}

func RenderJson(c *gin.Context, data interface{}) {
	//saveRespLog(c, data) //保存日志
	c.Header("Content-Type", "application/json;charset=UTF-8")
	c.JSON(200, data)
	c.Writer.Flush()
	c.Abort()
	return
}
