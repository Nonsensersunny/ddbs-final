package trace

import (
	"ddbs-final/internal/config"
	"ddbs-final/internal/utils"
	"ddbs-final/pkg/db/mysql"
	"ddbs-final/pkg/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateTrace(c *gin.Context) {
	trace := model.Trace{}
	if err := c.ShouldBindJSON(&trace); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	fmt.Println(trace)
	if trace.Current == "" || trace.Oid == "" {
		c.JSON(http.StatusBadRequest, utils.ErrorHelper(errors.New("bad request"), utils.PARSE_PARAMETER_ERROR))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(trace.Current)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_TRACE_FAIL))
		return
	}
	if err := trace.CreateTrace(mysql.GetMySQLDB(dbConf)); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_TRACE_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func GetTrace(c *gin.Context) {
	oid := c.Query("oid")
	var traces []*model.Trace
	for _, db := range config.GetServerConf().MySQLDBs {
		trace, _ := model.GetByOid(mysql.GetMySQLDB(db.MySQLConfig), oid)
		if trace.Id != "" {
			traces = append(traces, trace)
		}
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("traces", traces)))
}

func UpdateTrace(c *gin.Context) {
	area, err := c.Cookie("area")
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.AREA_INFO_LOST))
		return
	}

	dbConf, err := config.GetMySQLConfigByArea(area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.UPDATE_TRACE_FAIL))
		return
	}

	tid := c.Query("id")
	next := c.Query("next")
	trace, err := model.GetTraceById(mysql.GetMySQLDB(dbConf), tid)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.UPDATE_TRACE_FAIL))
		return
	}
	if err := trace.UpdateNext(mysql.GetMySQLDB(dbConf), next); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.UPDATE_TRACE_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}