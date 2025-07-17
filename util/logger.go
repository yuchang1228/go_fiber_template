package util

import (
	"os"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func SetupLog() {
	const logPath = "./logs/fiber.log"

	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", os.ModePerm)
	}

	os.OpenFile(logPath, os.O_RDONLY|os.O_CREATE, 0666)
	c := zap.NewProductionConfig()
	c.OutputPaths = []string{"stdout", logPath}
	Logger, _ = c.Build()
}
