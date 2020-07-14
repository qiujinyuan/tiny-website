package util

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	"github.com/yrjkqq/tiny-website/pkg/gredis"
	"github.com/yrjkqq/tiny-website/pkg/setting"
)

// Claims ...
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// TokenDetails ...
type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUUID   string
	RefreshUUID  string
	AtExpires    int64
	RtExpires    int64
}

// AccessDetails ...
type AccessDetails struct {
	AccessUUID string
	UserID     string
}

// GenerateToken generate a token
func GenerateToken(userID string) (td *TokenDetails, err error) {
	td = &TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * setting.AccessTokenExpires).Unix()
	td.AccessUUID = uuid.NewV4().String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * setting.RefreshTokenExpires).Unix()
	td.RefreshUUID = uuid.NewV4().String()

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUUID
	atClaims["user_id"] = userID
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(setting.AccessJwtSecret))
	if err != nil {
		return
	}

	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_id"] = td.RefreshUUID
	rtClaims["user_id"] = userID
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(setting.RefreshJwtSecret))
	return
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

// VerifyToken verify whether the token is valid
func VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(setting.AccessJwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

// TokenValid is the token valid
func TokenValid(r *http.Request) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return fmt.Errorf("token is invalid")
	}

	return nil
}

// ExtractTokenMetadata ...
func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUUID, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("access_uuid not found")
		}
		userID, ok := claims["user_id"].(string)
		if !ok {
			return nil, fmt.Errorf("user_id not found")
		}
		return &AccessDetails{accessUUID, userID}, nil
	}
	return nil, fmt.Errorf("token is invalid")
}

// FetchAuth fetch user in redis
func FetchAuth(ad *AccessDetails) (userID string, err error) {
	userID, err = gredis.Get(ad.AccessUUID)
	return
}
