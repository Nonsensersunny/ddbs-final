package user

import (
	"ddbs-final/internal/config"
	"ddbs-final/internal/utils"
	"ddbs-final/pkg/db/mysql"
	"ddbs-final/pkg/db/redis"
	"ddbs-final/pkg/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"time"
)

func UserRegister(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	if ok := utils.PhoneCheck(user.Phone); !ok {
		c.JSON(http.StatusOK, utils.ErrorHelper(errors.New("invalid phone number"), utils.INVALID_PHONE))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(user.Area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.INVALID_AREA))
		return
	}
	if ok := model.UserExistsWithPhone(mysql.GetMySQLDB(dbConf), user.Phone); ok {
		c.JSON(http.StatusOK, utils.ErrorHelper(errors.New("phone conflicts"), utils.USER_ALREADY_EXISTS))
		return
	}
	if err := user.CreateUser(mysql.GetMySQLDB(dbConf)); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.CREATE_USER_FAIL))
		return
	}
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func UserLogin(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.PARSE_PARAMETER_ERROR))
		return
	}
	dbConf, err := config.GetMySQLConfigByArea(user.Area)
	if err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.USER_LOGIN_FAIL))
		return
	}
	u, err := model.GetUserByPhoneAndPwd(mysql.GetMySQLDB(dbConf), user.Phone, user.Password)
	if err != nil || u.Id == "" {
		c.JSON(http.StatusOK, utils.ErrorHelper(errors.New("user login failed"), utils.USER_LOGIN_FAIL))
		return
	}
	cookieStr := uuid.NewV4().String()
	redisDB := redis.GetRedisClient(config.GetRedisConfig())
	if err := redisDB.Set(u.Id, cookieStr, time.Hour*12).Err(); err != nil {
		c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.SET_COOKIE_FAIL))
		return
	}
	c.SetCookie("uid", u.Id, 3600, "/", "", false, false)
	c.SetCookie("area", u.Area, 3600, "/", "", false, false)
	c.SetCookie("role", fmt.Sprintf("%d", u.Role), 3600, "/", "", false, false)
	c.SetCookie("token", cookieStr, 3600, "/", "", false, false)
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}

func UserLogout(c *gin.Context) {
	uid, _ := c.Cookie("uid")
	redisDB := redis.GetRedisClient(config.GetRedisConfig())
	redisDB.Del(uid)
	c.JSON(http.StatusOK, utils.RespHelper(utils.SuccessResp()))
}