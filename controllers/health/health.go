package health

import (
	"destijatim/configs"
	"destijatim/models/base"
	"destijatim/models/health"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckHealth(c echo.Context) error {
	conn, err := configs.DB.DB()
	if err != nil {
		fmt.Println("error instancing db: ", err.Error())
		return c.JSON(http.StatusInternalServerError,
			base.Response{
				Code:    500,
				Status:  "error",
				Message: err.Error(),
			})
	}
	err = conn.Ping()
	if err != nil {
		fmt.Println("error pinging db: ", err)
		return c.JSON(http.StatusInternalServerError,
			base.Response{
				Code:    500,
				Status:  "error",
				Message: err.Error(),
			})
	}
	return c.JSON(http.StatusInternalServerError,
		base.Response{
			Code:    200,
			Status:  "success",
			Message: "nothing wrong with DB",
			Data: health.Health{
				Status: "UP",
			},
		})
}
