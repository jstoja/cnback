package db

import (
	"fmt"
	"github.com/jstoja/cnback/config"
	"github.com/Sirupsen/logrus"
	"io"
	"os/exec"
)

func backupMongodb(config config.MongoDB) (*exec.Cmd, io.ReadCloser, io.ReadCloser, error) {
	args := fmt.Sprintf("--archive --gzip --host %v --port %v ", config.Host, config.Port)
	if config.Database != "" {
		args += fmt.Sprintf("--db %v ", config.Database)
	}
	if config.Username != "" && config.Password != "" {
		args += fmt.Sprintf("-u %v -p %v ", config.Username, config.Password)
	}
	if config.Params != "" {
		args += fmt.Sprintf("%v", config.Params)
	}

  logrus.Infof("Goging to launch mongodump %v", args)
	cmd := exec.Command("sh", "-c", "mongodump", args)


  out2, err := cmd.StderrPipe()
	if err != nil {
		logrus.Warn(err)
		return cmd, nil, out2, err
	}
	out1, err := cmd.StdoutPipe()
	if err != nil {
		logrus.Warn(err)
		return cmd, out1, out2, err
	}
	if err := cmd.Start(); err != nil {
		logrus.Fatal(err)
	}

	return cmd, out1, out2, nil
}
