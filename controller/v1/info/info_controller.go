package info

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"com.wangzhumo.distribute/conf"
	"com.wangzhumo.distribute/database/mysql"
	"com.wangzhumo.distribute/models"
	"com.wangzhumo.distribute/models/model"
	"github.com/gin-gonic/gin"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("info")
	rr.GET("/list", GetApkVersionList)
	rr.GET("/apkinfo", GetApkDetailsInfo)
	rr.GET("/version", GetApkListByVersion)
	rr.POST("/mark", MarkApkInfo)
}

func GetApkVersionList(c *gin.Context) {
	lastId := c.Query("last_id")
	size := c.Query("size")
	typeValue := c.Query("type")
	if lastId == "" || size == "" || typeValue == "" {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "参数异常",
		})
	}
	// 开始查询
	sizeInt, err := strconv.Atoi(size)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
	}
	if sizeInt == 0 {
		sizeInt = 10
	}
	lastIdInt, err := strconv.Atoi(lastId)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
	}

	if typeValue == "" {
		typeValue = "all"
	}
	listData := []model.Version{}
	// 判断应该查询什么类型的
	if typeValue == "release" {
		mysql.InstanceDB().Where("release_version = ? AND id > ?", true, lastIdInt).Limit(sizeInt).Find(&listData)
	} else if typeValue == "debug" {
		mysql.InstanceDB().Where("release_version = ? AND id > ?", false, lastIdInt).Limit(sizeInt).Find(&listData)
	} else {
		mysql.InstanceDB().Where("id > ?", lastIdInt).Limit(sizeInt).Find(&listData)
	}
	resp := models.VsersionListResponse{}
	if len(listData) == 0 || len(listData) < sizeInt {
		fillQrcodeToResponse(listData)
		resp.LastId = 0
		resp.More = false
		resp.Data = listData
	} else {
		fillQrcodeToResponse(listData)
		resp.LastId = int64(listData[len(listData)-1].ID)
		resp.More = true
		resp.Data = listData
	}
	c.JSON(http.StatusOK, &resp)
}

func fillQrcodeToResponse(listData []model.Version) {
	if len(listData) > 0 {
		for i, v := range listData {
			v.QrcodeURL = fmt.Sprintf("%s=%d", "v1/apk/download?id", v.ID)
			v.ApkURL = strings.Replace(v.ApkURL,".data","static",1)
			listData[i] = v
		}
	}
}

func GetApkListByVersion(c *gin.Context) {

}

func GetApkDetailsInfo(c *gin.Context) {
	id := c.Query("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
	}
	// 查询数据即可
	version := model.Version{}
	mysql.InstanceDB().First(&version, idInt)
	if version.ID == 0 {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_NOT_FOUND,
			Message: "找不到该数据",
		})
	}

	// 返回查询结果
	version.QrcodeURL = fmt.Sprintf("%s=%d", "v1/apk/download?id", version.ID)
	models.Success(c, &version)
}

func MarkApkInfo(c *gin.Context) {

}
