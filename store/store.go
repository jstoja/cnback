package store

import (
  "io"
)

type Store interface {
  io.WriteCloser
}

func SendBackup(backupStream io.ReadCloser) error {
	return nil
}
