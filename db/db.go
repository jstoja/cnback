package db

import (
  "io"
	"github.com/pkg/errors"
	"github.com/jstoja/cnback/config"
)

func FetchBackup(plan config.Plan) (io.ReadCloser, error) {
  if plan.MongoDB != nil {
    backupStream, _, err := backupMongodb(*plan.MongoDB)
    if err != nil {
      return nil, errors.Wrapf(err, "issue backuping mongodb")
    }
    // TODO: Handle logstream
    return backupStream, nil
  }
	return nil, nil
}
