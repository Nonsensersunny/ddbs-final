package middlewares

import (
	"ddbs-final/internal/config"
	"ddbs-final/internal/utils"
	"ddbs-final/pkg/db/redis"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COOKIE_NOT_FOUND))
			return
		}
		token, err := c.Cookie("token")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COOKIE_NOT_FOUND))
			return
		}
		redisToken, _ := redis.GetRedisClient(config.GetRedisConfig()).Get(uid).Result()
		if redisToken == token {
			c.Next()
			return
		}
		c.Abort()
		c.JSON(http.StatusUnauthorized, utils.ErrorHelper(errors.New("user not login"), utils.USER_NOT_LOGIN))
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := c.Cookie("uid")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COOKIE_NOT_FOUND))
			return
		}
		role, err := c.Cookie("role")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COOKIE_NOT_FOUND))
			return
		}
		if role != "2" {
			c.Abort()
			c.JSON(http.StatusUnauthorized, utils.ErrorHelper(err, utils.NOT_ADMIN))
			return
		}
		token, err := c.Cookie("token")
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, utils.ErrorHelper(err, utils.COOKIE_NOT_FOUND))
			return
		}
		redisToken, _ := redis.GetRedisClient(config.GetRedisConfig()).Get(uid).Result()
		if redisToken == token {
			c.Next()
			return
		}
		c.Abort()
		c.JSON(http.StatusUnauthorized, utils.ErrorHelper(errors.New("user not login"), utils.USER_NOT_LOGIN))
	}
}