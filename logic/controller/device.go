package controller

import (
	"goim/logic/entity"

	"goim/logic/service"

	"github.com/gin-gonic/gin"
)

func init() {
	g := Engine.Group("/device")
	g.POST("", DeviceController{}.Regist)
}

type DeviceController struct{}

// Regist 设备注册
func (DeviceController) Regist(c *gin.Context) {
	var device entity.Device
	err := c.ShouldBindJSON(&device)
	if err != nil {
		c.JSON(OK, BadRequest)
		return
	}
	id, token, err := service.NewDeviceService().Regist(device)
	if err != nil {
		c.JSON(OK, InternalServerError)
		return
	}

	result := make(map[string]interface{}, 2)
	result["id"] = id
	result["token"] = token

	c.JSON(OK, result)

}
