package models

import (
	"net/http"

	"com.wangzhumo.distribute/conf"
	"com.wangzhumo.distribute/models/model"
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ListResponse struct {
	LastId int64         `json:"last_id"`
	More   bool          `json:"more"`
	Data   []interface{} `json:"data"`
}

type VsersionListResponse struct {
	LastId int64           `json:"last_id"`
	More   bool            `json:"more"`
	Data   []model.Version `json:"data"`
}

//请求成功的时候 使用该方法返回信息
func Success(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    conf.STATUS_OK,
		"message": "ok",
		"data":    v,
	})
}

//请求成功的时候 使用该方法返回信息
func SuccessList(ctx *gin.Context, more bool, lastId int64, v []interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"last_id": lastId,
		"more":    more,
		"data":    v,
	})
}

//请求失败的时候, 使用该方法返回信息
func Failed(ctx *gin.Context, v interface{}) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code":    conf.STATUS_FAILED,
		"data":    nil,
		"message": v,
	})
}
