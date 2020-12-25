package http

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go-framework/http/pkg/common"
	"go-framework/http/pkg/logger"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
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

// HttpPost post请求
func PostJson(url string, params []byte) ([]byte, error) {
	body := bytes.NewBuffer(params)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("httpcode error:" + fmt.Sprint(resp.StatusCode))
	}

	return respData, nil
}

func PostJsonRequest(url string, requestData interface{}, responseData interface{}) (err error) {
	if !common.IsValidString(url) {
		return errors.New("PostJsonRequest url is invalid!")
	}

	var jsonData []byte
	jsonData, err = json.Marshal(requestData)
	if err != nil {
		logger.Warnf("Marshal json error:%s ", err.Error())
		return err
	}
	log.Printf("request:%+v, json:%+v\n", requestData, string(jsonData))
	var responseByte []byte
	responseByte, err = PostJson(url, jsonData)
	if err != nil {
		logger.Warnf("PostJson error:%s url:%s data:%s", err.Error(), url, string(responseByte))
		return err
	}
	responseByte = []byte(strings.Replace(string(responseByte), "\"data\":[]", "\"data\":{}", 1))
	err = json.Unmarshal(responseByte, responseData)
	log.Printf("url:%s ,request:%+v, response:%+v\n", url, requestData, string(responseByte))
	if err != nil {
		logger.Warnf("Unmarshal response error,data:%s err:%s ", string(responseByte), err.Error())
		return err
	}
	return nil
}
