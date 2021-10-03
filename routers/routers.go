package routers

import (
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

// InitRoute initiate Route
func InitRoute() {
	port := viper.Get("PORT").(int)
	e := echo.New()
	err := e.Start(":" + strconv.Itoa(port))
	if err != nil {
		panic("cannot start service: " + err.Error())
	}
}
