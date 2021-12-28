package apk

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"com.wangzhumo.distribute/conf"
	"com.wangzhumo.distribute/database/mysql"
	"com.wangzhumo.distribute/models"
	"com.wangzhumo.distribute/models/model"
	"com.wangzhumo.distribute/utils/parser"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func Routers(r *gin.RouterGroup) {
	rr := r.Group("apk")
	rr.POST("/create", createApkProject)
	rr.POST("/upload", uploadApkFile)
	rr.GET("/download", downloadApkFile)
}

// 创建一个APK项目
func createApkProject(c *gin.Context) {
	//apkname := c.PostForm("name")
	iconfile, err := c.FormFile("icon")
	name := c.PostForm("name")
	appid := c.PostForm("appid")

	findApk := &model.Apkinfo{ID: 0}
	// 先查找
	mysql.InstanceDB().Where(&model.Apkinfo{
		AppID: appid,
	}).Find(findApk)
	if findApk.ID != 0 {
		// 否则，是重复创建
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.STATUS_OK,
			Data:    findApk,
			Message: "不能重复创建",
		})
		return
	}

	// params
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
		return
	}
	ext := models.BundleFileExtension(filepath.Ext(iconfile.Filename))
	if !ext.IsValidImage() {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "icon image error",
		})
		return
	}

	// 保存icon地址
	_uuid := uuid.NewV4().String()
	filename := filepath.Join(".data/icons", _uuid+string(ext))
	err = c.SaveUploadedFile(iconfile, filename)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
		return
	}
	// 写入数据库
	info := model.Apkinfo{
		ApkIcon:     filename,
		ApkName:     name,
		AppID:       appid,
		LastRelease: 0,
	}

	ret := mysql.InstanceDB().Create(&info)
	if ret.Error != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: ret.Error.Error(),
		})
		return
	}

	// 否则，创建成功
	c.JSON(http.StatusOK, &models.Response{
		Code:    conf.STATUS_OK,
		Message: "创建成功",
		Data:    info,
	})
}

// 上传Apk文件
func uploadApkFile(c *gin.Context) {
	// 1.收到文件
	apkFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
		return
	}
	release := c.PostForm("type")
	changelog := c.PostForm("changelog")
	ext := models.BundleFileExtension(filepath.Ext(apkFile.Filename))
	if !ext.IsValidApk() {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "apk file error",
		})
		return
	}

	// 保存apk地址
	_, err = os.Stat(".data/temp")
	if os.IsNotExist(err) {
		os.MkdirAll(".data/temp", 0755)
	}

	_uuid := uuid.NewV4().String()
	filename := filepath.Join(".data/temp", _uuid+string(ext))
	err = c.SaveUploadedFile(apkFile, filename)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
		return
	}

	// 2.解析文件
	apkInfo, err := parser.NewAppParser(filename)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: err.Error(),
		})
		return
	}
	findApk := &model.Apkinfo{ID: 0}
	mysql.InstanceDB().Where(&model.Apkinfo{
		AppID: apkInfo.BundleId,
	}).Find(findApk)

	versionCode, _ := strconv.Atoi(apkInfo.Build)
	version := &model.Version{
		Name:      apkInfo.Version,
		Code:      int32(versionCode),
		Downloads: 0,
		Changelog: changelog,
		ApkID:     findApk.ID,
		Timestamp: time.Now().UnixMilli(),
	}

	// 3.写入文件
	dstPath := fmt.Sprintf(".data/apks/%s", apkInfo.Build)
	_, err = os.Stat(dstPath)
	if os.IsNotExist(err) {
		os.MkdirAll(dstPath, 0755)
		os.RemoveAll(dstPath)
	}
	dstName := filepath.Join(dstPath, apkFile.Filename)
	os.Rename(filename, dstName)
	os.Remove(filename)

	// 3.保存到数据库
	version.ReleaseVersion = release == "release"
	version.ApkURL = dstName
	ret := mysql.InstanceDB().Create(&version)
	if ret.Error != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: ret.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, &models.Response{
		Code:    conf.STATUS_OK,
		Message: "ok",
		Data:    version,
	})
}

// 下载Apk文件
func downloadApkFile(c *gin.Context) {
	versionId, err := strconv.Atoi(c.Query("id"))
	if err != nil || versionId == 0 {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "需要AppID参数",
		})
		return
	}

	// 获取ID
	findApk := &model.Version{ID: 0}
	// 先查找
	mysql.InstanceDB().Where(&model.Version{
		ID: int32(versionId),
	}).Find(findApk)

	if findApk.ID == 0 {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "找不到该版本",
		})
	}

	// 写入qrcode
	findApk.QrcodeURL = fmt.Sprintf("%s=%d", "v1/apk/download?id", findApk.ID)
	file, err := os.OpenFile(findApk.ApkURL, 2, 0666)
	if err != nil {
		c.JSON(http.StatusOK, &models.Response{
			Code:    conf.ERROR_PARAMS_INVIDE,
			Message: "打开文件失败",
		})
		return
	}
	defer file.Close()
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name()))
	c.Writer.Header().Set("Content-Type", "application/zip")
	// 去获取文件返回
	c.File(findApk.ApkURL)
}
