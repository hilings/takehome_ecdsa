package util

import (
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

// LoadRequestBody ...
func LoadRequestBody(c echo.Context, req interface{}) {
	bytes, err := ioutil.ReadAll(c.Request().Body)
	defer c.Request().Body.Close()
	if err != nil {
		c.Logger().Warnf("ReadAll failed, error: %v, body: %+v", err, c.Request().Body)
		return
	}

	err = json.Unmarshal(bytes, &req)
	if err != nil {
		c.Logger().Warnf("Unmarshal failed, error: %v, bytes: %s", err, string(bytes))
	}

	c.Logger().Infof("Request: %+v", req)
}

// LoadRequestHeader ...
func LoadRequestHeader(c echo.Context, key string) string {
	header := c.Request().Header
	values := header[key]
	if len(values) == 0 {
		return ""
	}

	return values[0]
}
