package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func AccessTokenValidation(tokenString, accessUuid, userId string) error {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New(fmt.Sprintf("unexpected signing method: %v", token.Header["alg"]))
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if res, err := checkTokenExpires(claims, accessUuid, userId); res {
				return []byte(os.Getenv("ACCESS_SECRET")), nil
			} else {
				return nil, err
			}
		} else {
			return nil, errors.New("claims check error")
		}
	})
	if err != nil {
		return err
	}
	return nil
}
func checkTokenExpires(claims jwt.MapClaims, accessUuid, userId string) (bool, error) {
	curTime := time.Now().Unix()
	var tokenTime time.Time

	exp, ok := claims["exp"].(float64)
	if !ok {
		return false, errors.New("no expiration meta")
	}
	tokenTime = time.Unix(int64(exp), 0)
	if curTime > tokenTime.Unix() {
		return false, errors.New("token expired")
	}

	accessUuidInToken, ok := claims["access_uuid"].(string)
	if !ok {
		return false, errors.New("no access uuid")
	}
	if accessUuid != accessUuidInToken {
		return false, errors.New("wrong access uuid")
	}

	userIdInToken, ok := claims["user_id"].(string)
	if !ok {
		return false, errors.New("no user id")
	}
	if userId != userIdInToken {
		return false, errors.New("wrong user id")
	}

	return true, nil
}

//func ExtractTokenMetadata(r *http.Request) (*AccessDetails, error) {
//	token, err := verifyToken(r)
//	if err != nil {
//		return nil, err
//	}
//	claims, ok := token.Claims.(jwt.MapClaims)
//	if ok && token.Valid {
//		accessUuid, ok := claims["access_uuid"].(string)
//		if !ok {
//			return nil, err
//		}
//		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
//		if err != nil {
//			return nil, err
//		}
//		return &AccessDetails{
//			AccessUuid: accessUuid,
//			UserId:     int64(userId),
//		}, nil
//	}
//	return nil, err
//}
//func TokenAuthMiddleware() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		err := tokenValid(c.Request)
//		if err != nil {
//			c.JSON(http.StatusUnauthorized, err.Error())
//			c.Abort()
//			return
//		}
//		c.Next()
//	}
//}

//func extractToken(r *http.Request) (string, error) {
//	bearToken := r.Header.Get("Authorization")
//	strArr := strings.Split(bearToken, " ")
//	if len(strArr) == 2 {
//		return strArr[1], nil
//	}
//	return "", errors.New("Token not given")
//}
