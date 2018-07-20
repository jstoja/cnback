package db

import (
  "io"
	"os"
	"github.com/pkg/errors"
	"github.com/Sirupsen/logrus"
	"github.com/jstoja/cnback/config"
)

func logErrors(plan config.Plan, errorStream io.ReadCloser) {
  // read from the PipeReader to stdout
  if _, err := io.Copy(os.Stdout, errorStream); err != nil {
    logrus.Fatalf("cannot read stderr from command: %v", err)
  }
}

func FetchBackup(plan config.Plan) (io.ReadCloser, error) {
  if plan.MongoDB != nil {
    cmd, backupStream, errorStream, err := backupMongodb(*plan.MongoDB)
    logErrors(plan, errorStream)

    // Channel with error + timer?
    if err := cmd.Wait(); err != nil {
      logrus.Fatalf("command failed: %v", err)
	  }

    // 1. Find how to wrap this nicely
    //    passing cmd around is not a good idea
    // 2. integrate logs from commands correctly


    logrus.Infof("finished command: %v", cmd.Args)
    if err != nil {
      return nil, errors.Wrapf(err, "issue backuping mongodb")
    }
    return backupStream, nil
  }
  logrus.Info("Didn't find any relevant source to backup")
	return nil, nil
}
