package middleware

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func FancyLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				fmt.Printf("%s [ERROR] %s\n    Path: %s\n    Query: %s\n    Error: %s\n",
					color.RedString(end.Format("2006/01/02 15:04:05")),
					color.RedString("🚨 Request error"),
					color.CyanString(path),
					color.CyanString(query),
					color.RedString(e),
				)
			}
		} else {
			statusCode := c.Writer.Status()
			var icon string
			var statusColor func(a ...interface{}) string
			switch {
			case statusCode >= 500:
				icon = "💥"
				statusColor = color.New(color.FgRed).SprintFunc()
			case statusCode >= 400:
				icon = "⚠️"
				statusColor = color.New(color.FgYellow).SprintFunc()
			case statusCode >= 300:
				icon = "↪️"
				statusColor = color.New(color.FgCyan).SprintFunc()
			default:
				icon = "✅"
				statusColor = color.New(color.FgGreen).SprintFunc()
			}

			fmt.Printf("%s [INFO] %s %s\n    Method: %s\n    Path: %s\n    Query: %s\n    Status: %s\n    Latency: %s\n    IP: %s\n    User-Agent: %s\n",
				color.BlueString(end.Format("2006/01/02 15:04:05")),
				icon,
				color.MagentaString("Request completed"),
				color.CyanString(c.Request.Method),
				color.CyanString(path),
				color.CyanString(query),
				statusColor(fmt.Sprintf("%d", statusCode)),
				color.YellowString(latency.String()),
				color.CyanString(c.ClientIP()),
				color.CyanString(c.Request.UserAgent()),
			)
		}
	}
}
