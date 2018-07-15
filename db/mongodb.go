package db

import (
	"fmt"
	"github.com/jstoja/cnback/config"
	"io"
	"os/exec"
)

func backupMongodb(config config.MongoDB) (io.ReadCloser, io.ReadCloser, error) {
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

	cmd := exec.Command("mongodump", args)
	cmd.Start()

	out2, err := cmd.StderrPipe()
	if err != nil {
		return nil, out2, err
	}
	out1, err := cmd.StdoutPipe()
	if err != nil {
		return out1, out2, err
	}

	return out1, out2, nil
}
