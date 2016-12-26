package main

import (
"encoding/json"
"github.com/labstack/echo"
"net/http"
)

func Postlog(c echo.Context) error {
	result := c.FormValue(Opts.Param)
	if result == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	c.Set("logMessage", result)
	return c.Blob(http.StatusOK, "image/gif", Gif)
}

func Getlog(c echo.Context) error {
	res := ""
	result := c.QueryParam(Opts.Param)
	if result != "" {
		res = result
	} else if r := c.QueryParams(); len(r) != 0 {
		req := c.Request()
		r["referer"] = []string{req.Referer()}
		r["UA"] = []string{req.UserAgent()}
		resByte, err := json.Marshal(r)
		if err != nil {
			return c.NoContent(http.StatusBadRequest)
		}
		res = string(resByte)
	}
	if res == "" {
		return c.NoContent(http.StatusBadRequest)
	}
	c.Set("logMessage", res)
	return c.Blob(http.StatusOK, "image/gif", Gif)
}

