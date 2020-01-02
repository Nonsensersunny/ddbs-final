package utils

import "github.com/gin-gonic/gin"

func ErrorHelper(err error, statusCode int) gin.H {
	return gin.H{
		"error": err,
		"code":  statusCode,
		"data":  "",
	}
}

type M map[string]interface{}

func SetData(key string, val interface{}) M {
	return M{
		key: val,
	}
}

func SuccessResp() M {
	return SetData("status", "success")
}

func RespHelper(m M) gin.H {
	return gin.H{
		"code":  SUCCESS,
		"data":  m,
		"error": "",
	}
}