package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/yrjkqq/tiny-website/models"
	"github.com/yrjkqq/tiny-website/pkg/e"
	"github.com/yrjkqq/tiny-website/pkg/gredis"
	"github.com/yrjkqq/tiny-website/pkg/util"
)

type auth struct {
	Username string `valid:"Required; MaxSize(50)"`
	Password string `valid:"Required; MaxSize(50)"`
}

func createAuth(userID string, td *util.TokenDetails) error {
	at := time.Unix(td.AtExpires, 0)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := gredis.Set(td.AccessUUID, userID, at.Sub(now))
	if errAccess != nil {
		return errAccess
	}

	errRefresh := gredis.Set(td.RefreshUUID, userID, rt.Sub(now))
	if errRefresh != nil {
		return errRefresh
	}

	return nil
}

// GetAuth login and get token
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})
	code := e.InvalidParams
	var msg string

	if ok {
		isExist, auth := models.CheckAuth(username, password)
		if isExist {
			td, err := util.GenerateToken(auth.ID.String())
			if err != nil {
				code = e.ErrorAuthToken
			} else {
				saveErr := createAuth(auth.ID.String(), td)
				if saveErr != nil {
					code = e.ErrorSaveToken
				} else {
					data["accessToken"] = td.AccessToken
					data["refreshToken"] = td.RefreshToken
					code = e.Success
				}
			}
		} else {
			code = e.ErrorAuth
		}
	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			msg += fmt.Sprintf(" %v: %v;", err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code) + msg,
		"data": data,
	})
}

func deleteAuth(givenUUID string) (bool, error) {
	return gredis.Del(givenUUID)
}

// Logout logout and delete token
func Logout(c *gin.Context) {
	var aID string
	accessUUID, exist := c.Get("accessUUID")
	if !exist {
		ad, err := util.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		aID = ad.AccessUUID
	} else {
		aID = accessUUID.(string)
	}
	ok, err := deleteAuth(aID)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, err.Error())
		return
	}
	if !ok {
		c.JSON(http.StatusNotAcceptable, "logout failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": e.Success,
	})
}
