package middlewares

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	var ignoredPaths = map[string]bool{
		"/v1/files/":               true,
		"/v1/login":                true,
		"/v1/refresh_token":        true,
		"/dashboard/register":      true,
		"/dashboard/login":         true,
		"/dashboard/refresh_token": true,
		"/admin/login":             true,
		"/admin/refresh_token":     true,
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// record body
		var body []byte
		if !ignoredPaths[strings.ToLower(path)] && c.Request.ContentLength < 5*1024 {
			body, _ = ioutil.ReadAll(c.Request.Body)
			c.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		}

		// Process request
		c.Next()

		param := gin.LogFormatterParams{
			Request: c.Request,
			Keys:    c.Keys,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.String()

		param.BodySize = c.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		entry := logrus.WithFields(logrus.Fields{
			"status code":    param.StatusCode,
			"latency":        fmt.Sprintf("%13v", param.Latency),
			"client ip":      fmt.Sprintf("%15s", param.ClientIP),
			"method":         param.Method,
			"path":           param.Path,
			"full path":      c.FullPath(),
			"body":           string(body),
			"content length": c.Request.ContentLength,
		})

		if param.ErrorMessage != "" {
			entry = entry.
				WithField("errors", param.ErrorMessage).
				WithField("stack", c.GetString("error_stack"))

			if c.GetString("error_stack") != "" {
				fmt.Printf("Request error %v\n%v\n", param.ErrorMessage, c.GetString("error_stack"))
			}
		}
		entry.Info("Request")
	}
}
