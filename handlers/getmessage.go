package handlers

import (
	"myapp/common"
	"myapp/util"
	"net/http"

	"github.com/labstack/echo"
)

// GetMessageResponse ...
type GetMessageResponse struct {
	Message string `json:"message"`
}

// GetMessage ...
func GetMessage(c echo.Context) error {
	publicKey := util.LoadRequestHeader(c, common.PublicKey)
	if len(publicKey) == 0 {
		c.Logger().Warn("empty publicKey")
		return c.JSON(http.StatusBadRequest, nil)
	}
	c.Logger().Debugf("publicKey: %s", publicKey)

	msg := util.RandomMessage(publicKey) // publicKey as seed for now
	c.Logger().Debugf("msg: %s", msg)

	resp := &GetMessageResponse{
		Message: msg,
	}
	return c.JSON(http.StatusOK, resp)
}
