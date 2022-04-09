package main

import (
	"github.com/quantumsheep/plouf"
	"github.com/quantumsheep/plouf/example/modules"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	worker, err := plouf.NewWorker(&modules.MainModule{})
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Fatal(worker.Start("localhost:8080"))
}
