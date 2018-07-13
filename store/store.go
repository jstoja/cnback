package store

import "io"

type Store interface {
  io.WriteCloser
}
