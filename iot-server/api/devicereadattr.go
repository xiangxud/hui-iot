package api

/**
 * @Author: lgh-dev
 * @Author: lgh-inter@163.com
 * @Date: 2020/1/9 21:39
 * @Desc:
 */
import (
	"github.com/gin-gonic/gin"
	"hui-iot/iot-server/common"
	"net/http"
)

func UpdateDeviceReadAttr(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}

func GetDeviceReadAttr(c *gin.Context) {
	c.JSON(http.StatusOK, common.BuildSucc(&common.ResultDTO{}))
}