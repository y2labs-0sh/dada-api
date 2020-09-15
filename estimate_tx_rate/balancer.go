package estimatetxrate

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
)

// 0x9B208194Acc0a8cCB2A8dcafEACfbB7dCc093F81

func balancerHandler(c echo.Context) error {

	return c.JSON(http.StatusInternalServerError, errors.New("Not impl yet"))
}
