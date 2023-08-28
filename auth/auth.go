package auth

import (
	"crypto/rand"
	"encoding/hex"
	"myapp/service"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/markbates/goth/gothic"
)

// set to true for dev mode(api allowing all calls)
var devMode = false

func HandleAuthCallback(c echo.Context) error {
	user, err := gothic.CompleteUserAuth(c.Response(), c.Request())

	if err != nil {
		return c.JSON(http.StatusForbidden, err)
	}

	userWithSameId := service.GetUser(user.UserID)
	userWithSameId.Token = GenerateSecureToken(30)
	timeInHalfAnHour := time.Now().Add(30 * time.Minute)
	userWithSameId.TokenExpiration = strconv.FormatInt(int64(timeInHalfAnHour.Unix()), 10)

	if userWithSameId.ID == "" {
		userWithSameId.ID = user.UserID
		userWithSameId.Username = user.NickName
		service.SaveUser(userWithSameId)
	} else {
		service.UpdateUser(userWithSameId)
	}

	tokenCookie := http.Cookie{
		Name:  "TOKEN",
		Value: userWithSameId.Token,
		Path:  "/",
	}

	userNameCookie := http.Cookie{
		Name:  "USERNAME",
		Value: user.NickName,
		Path:  "/",
	}

	c.SetCookie(&tokenCookie)
	c.SetCookie(&userNameCookie)
	return c.Redirect(http.StatusPermanentRedirect, "http://localhost:3000/products/read")
}

func HandleAuth(c echo.Context) error {
	c.Request().URL.RawQuery = "provider=google"
	gothic.BeginAuthHandler(c.Response(), c.Request())
	return c.JSON(http.StatusOK, c.Request())
}

func Logout(c echo.Context) error {
	tokenCookie, err := c.Request().Cookie("TOKEN")
	if err != nil {
		println(err.Error())
		return echo.NewHTTPError(http.StatusUnauthorized, "You are not logged in")
	}
	user := service.GetUserByToken(tokenCookie.Value)
	user.Token = ""
	user.TokenExpiration = "0"
	service.UpdateUser(user)
	newTokenCookie := http.Cookie{
		Name:  "TOKEN",
		Value: "",
		Path:  "/",
	}

	userNameCookie := http.Cookie{
		Name:  "USERNAME",
		Value: "",
		Path:  "/",
	}

	c.SetCookie(&newTokenCookie)
	c.SetCookie(&userNameCookie)
	return c.String(http.StatusOK, "Logged out")

}
func GenerateSecureToken(length int) string {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return hex.EncodeToString(b)
}

func AuthHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if devMode {
			return next(c)
		}
		if strings.Contains(c.Request().RequestURI, "auth") {
			return next(c)
		}
		tokenCookie, err := c.Request().Cookie("TOKEN")
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Your token has expired")
		}
		user := service.GetUserByToken(tokenCookie.Value)

		if user.ID != "" {
			tm, err := strconv.ParseInt(user.TokenExpiration, 10, 64)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Token invalid")
			}
			tmNow := int64(time.Now().Unix())
			if tm-tmNow > 0 {
				return next(c)
			}
		}
		return echo.NewHTTPError(http.StatusUnauthorized, "Your token has expired")
	}
}
