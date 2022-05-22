package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

// GetMessageResponse ...
type GetMessageResponse struct {
	Message string `json:"message"`
}

// GetMessage ...
func GetMessage(c echo.Context) error {

	// TODO
	resp := &GetMessageResponse{
		Message: "random_message",
	}

	return c.JSON(http.StatusNotImplemented, resp)
}
