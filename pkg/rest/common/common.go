package common

import (
	"ddbs-final/internal/config"
	"ddbs-final/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllAreas(c *gin.Context) {
	var areas []string
	for _, db := range config.GetServerConf().MySQLDBs {
		areas = append(areas, db.Area)
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SetData("areas", areas)))
}
