package main

import (
	"ddbs-final/internal/config"
	"ddbs-final/pkg/db/mysql"
	"ddbs-final/pkg/model"
	"ddbs-final/pkg/rest"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func startup() *config.ServerConf {
	serverConfig := config.GetServerConf()
	dbs := serverConfig.MySQLDBs
	for _, db := range dbs {
		dbClient := mysql.GetMySQLDB(db.MySQLConfig)
		mysql.MigrateTables(dbClient, &model.Order{}, &model.User{}, &model.Trace{})
	}
	return serverConfig
}

func main()  {
	serverConfig := startup()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = serverConfig.Http.AllowOrigin
	corsConfig.AllowCredentials = true
	r.Use(cors.New(corsConfig))

	rest.REST(r)
	r.Run(fmt.Sprintf("%s:%d", serverConfig.Http.Host, serverConfig.Http.Port))
}