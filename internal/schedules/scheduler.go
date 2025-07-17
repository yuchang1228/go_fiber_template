package main

import (
	"app/internal/jobs"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/go-co-op/gocron/v2"
)

var f *os.File

func initLogging() {
	logDir := "log"
	logFile := "scheduler.log"
	logPath := filepath.Join(logDir, logFile)

	file, _ := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	iw := io.MultiWriter(os.Stdout, file)
	log.SetOutput(iw)
}

func main() {
	initLogging()
	defer f.Close()

	s, err := gocron.NewScheduler()

	if err != nil {
		panic(err)
	}

	j, err := s.NewJob(
		gocron.CronJob("*/5 * * * * *", true),
		gocron.NewTask(
			jobs.PublishJob,
			"Test",
			"Hello World",
		),
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(j.ID())

	s.Start()

	select {}
}
