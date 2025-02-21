package middleware

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func SetupLogger() (*logrus.Logger, *os.File, error) {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err := os.Mkdir("logs", 0755)
		if err != nil {
			logrus.Fatal("Error creating logs folder: ", err)
		}
	}
	
	logFile, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		return nil, nil, err
	}

	logger := logrus.New()
	logger.SetOutput((io.MultiWriter(os.Stdout, logFile)))
	logger.SetLevel(logrus.DebugLevel)

	logger.SetFormatter(&logrus.JSONFormatter{})

	return logger, logFile, nil
}

func LoggerMiddleware(logger *logrus.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestID := fmt.Sprintf("%d", time.Now().UnixNano())

			startTime := time.Now()
			err := next(c)

			duration := time.Since(startTime)

			status := c.Response().Status

			var logLevel logrus.Level
			var logMessage string

			if status >= 200 && status < 300 {
				logLevel = logrus.InfoLevel
				logMessage = "Request processed successfully"
			} else if status >= 400 && status <= 500 {
				logLevel = logrus.WarnLevel
				logMessage = "Client error"
			} else if status >= 500 {
				logLevel = logrus.ErrorLevel
				logMessage = "Server error"
			}

			logFields := logrus.Fields{
				"method": c.Request().Method,
				"url": c.Request().URL.String(),
				"request_id": requestID,
				"duration": duration.Seconds(),
				"status": status,
			}

			if c.Get("error") != nil {
				logFields["error"] = c.Get("error")
			}

			logger.WithFields(logFields).Log(logLevel, logMessage)

			if err != nil {
				logger.WithFields(logrus.Fields{
					"method": c.Request().Method,
					"url": c.Request().URL.String(),
					"request_id": requestID,
					"error": err.Error(),
				}).Error("Request failed")
			}

			return err
		}
	}
}