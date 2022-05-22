package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// VerifyRequest ...
type VerifyRequest struct {
	Address       string `json:"address"`
	SignedMessage string `json:"signedMessage"`
}

// VerifyResponse ...
type VerifyResponse struct {
	Verified bool `json:"verified"`
}

// Verify ...
func Verify(c echo.Context) error {

	req := &VerifyRequest{}

	// TODO
	resp := &VerifyResponse{
		Verified: false,
	}

	bytes, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		c.Logger().Warnf("ReadAll failed, error: %v, body: %+v", err, c.Request().Body)
		return c.JSON(http.StatusBadRequest, resp)
	}

	err = json.Unmarshal(bytes, &req)
	if err != nil {
		c.Logger().Warnf("Unmarshal failed, error: %v, bytes: %s", err, string(bytes))
		return c.JSON(http.StatusBadRequest, resp)
	}

	c.Logger().Debugf("req = %+v\n", req)

	return c.JSON(http.StatusNotImplemented, resp)
}
