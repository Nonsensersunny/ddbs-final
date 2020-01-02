package order

import (
	"ddbs-final/internal/config"
	"ddbs-final/internal/utils"
	"ddbs-final/pkg/db/mysql"
	"ddbs-final/pkg/model"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreateOrder(c *gin.Context) {
	uid, _ := c.Cookie("uid")
	area, err := c.Cookie("area")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.AREA_INFO_LOST))
		return
	}
	var order model.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	order.Uid = uid
	dbConf, err := config.GetMySQLConfigByArea(area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_ORDER_FAILE))
		return
	}
	order.Finished = false
	_, err = order.CreateOrder(mysql.GetMySQLDB(dbConf))
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_ORDER_FAILE))
		return
	}

	startDB, err := config.GetMySQLConfigByArea(order.Start)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_ORDER_FAILE))
		return
	}
	startTrace := &model.Trace{
		Id:      "",
		Time:    time.Now(),
		Oid:     order.Id,
		Current: order.Start,
		Next:    "",
	}
	if err := startTrace.CreateTrace(mysql.GetMySQLDB(startDB)); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_ORDER_FAILE))
		return
	}

	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func GetOrder(c *gin.Context) {
	uid, _ := c.Cookie("uid")
	area, err := c.Cookie("area")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.AREA_INFO_LOST))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_USER_FAIL))
		return
	}
	id := c.Query("id")
	if id != "" {
		order, err := model.GetOrderById(mysql.GetMySQLDB(dbConf), id)
		if err != nil {
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.GET_ORDER_FAIL))
			return
		}
		c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("order", order)))
		return
	}
	orders, err := model.GetOrdersByUid(mysql.GetMySQLDB(dbConf), uid)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.GET_ORDER_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("orders", orders)))
	return
}

func GetAllOrders(c *gin.Context) {
	area, err := c.Cookie("area")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.AREA_INFO_LOST))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	orders, err := model.GetAllOrders(mysql.GetMySQLDB(dbConf))
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("orders", orders)))
}

func TerminateOrder(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(errors.New("id info lost"), utils.PARSE_PARAMETER_ERROR))
		return
	}
	area, err := c.Cookie("area")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.AREA_INFO_LOST))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_USER_FAIL))
		return
	}
	if err := model.FinishOrder(mysql.GetMySQLDB(dbConf), id); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.TERMINATE_ORDER_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}