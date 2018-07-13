package db

import (
	"fmt"
	"github.com/jstoja/cnback/config"
	"io"
	"os/exec"
)

func backup(plan config.Plan) (io.ReadCloser, io.ReadCloser, error) {
	dump := fmt.Sprintf("mongodump --archive --gzip --host %v --port %v ", plan.Source.Host, plan.Source.Port)
	if plan.Source.Database != "" {
		dump += fmt.Sprintf("--db %v ", plan.Source.Database)
	}
	if plan.Source.Username != "" && plan.Source.Password != "" {
		dump += fmt.Sprintf("-u %v -p %v ", plan.Source.Username, plan.Source.Password)
	}
	if plan.Source.Params != "" {
		dump += fmt.Sprintf("%v", plan.Source.Params)
	}

	cmd := exec.Command("/bin/sh", "-c", dump)
	cmd.Start()

	out2, err := cmd.StderrPipe()
	if err != nil {
		return nil, out2, err
	}
	out1, err := cmd.StdoutPipe()
	if err != nil {
		return nil, out1, err
	}

	return out1, out2, nil
}
