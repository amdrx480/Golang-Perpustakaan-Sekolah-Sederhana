package middlewares

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

// untuk menyimpan konfigurasi jwt
type JWTConfig struct {
	SecretKey       string
	ExpiresDuration int
}

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(jwtConfig.SecretKey),
	}
}

func (jwtConfig *JWTConfig) GenerateToken(adminId int) (string, error) {
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(int64(jwtConfig.ExpiresDuration))))

	claims := &JwtCustomClaims{
		adminId,
		jwt.RegisteredClaims{
			ExpiresAt: expire,
		},
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//mendapatkan token jwt dalam bentuk string
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))

	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(c echo.Context) (*JwtCustomClaims, error) {
	user := c.Get("user").(*jwt.Token)

	if user == nil {
		return nil, errors.New("invalid")
	}

	claims := user.Claims.(*JwtCustomClaims)

	return claims, nil
}

// menentukan token jwt valid atau tidak
func VerifyTooken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		adminData, err := GetUser(c)

		isInvalid := adminData == nil || err != nil

		if isInvalid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}
		return next(c)
	}
}
