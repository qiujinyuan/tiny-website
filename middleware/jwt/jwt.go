package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yrjkqq/tiny-website/pkg/util"
)

// JWT gin middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		ad, err := util.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}

		userID, err := util.FetchAuth(ad)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusUnauthorized, err.Error())
			return
		}
		c.Set("accessUUID", ad.AccessUUID)
		c.Set("userID", userID)
		c.Next()
	}
}
