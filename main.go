package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/Sirupsen/logrus"
	"github.com/jstoja/cnback/config"
	"github.com/jstoja/cnback/scheduler"
)

var version = "0.1.0"

func main() {
	var appConfig = &config.AppConfig{}
	flag.StringVar(&appConfig.LogLevel, "LogLevel", "debug", "logging threshold level: debug|info|warn|error|fatal|panic")
	flag.StringVar(&appConfig.ConfigPath, "ConfigPath", "/config", "plan yml files dir")
	flag.Parse()
	setLogLevel(appConfig.LogLevel)
	logrus.Infof("Starting with config: %+v", appConfig)

	//info, err := backup.CheckMongodump()
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Info(info)

	//info, err = backup.CheckMinioClient()
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Info(info)

	//info, err = backup.CheckGCloudClient()
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//logrus.Info(info)

	plans, err := config.LoadPlans(appConfig.ConfigPath)
	if err != nil {
		logrus.Fatal(err)
	}

	//store, err := db.Open(path.Join(appConfig.DataPath, "mgob.db"))
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	//statusStore, err := db.NewStatusStore(store)
	//if err != nil {
	//	logrus.Fatal(err)
	//}
	sch := scheduler.New(plans, appConfig, nil)
  scheduleCount, err := sch.Start()
  if err != nil {
    logrus.Errorf("Failed to launch a schedule: %v", err)
  }
  if scheduleCount == 0 {
    return
  }

	//wait for SIGINT (Ctrl+C) or SIGTERM (docker stop)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	sig := <-sigChan

	logrus.Infof("Shutting down %v signal received", sig)
}

func setLogLevel(levelName string) {
	level, err := logrus.ParseLevel(levelName)
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.SetLevel(level)
}
