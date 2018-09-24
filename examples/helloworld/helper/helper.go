package helper

import (
	"io"
)

// NewWriter turns sender into an io.Writer. The sender callback will
// receive []byte arguments of length at most WriteBufferSize.
func NewWriter(sender func(p []byte) error) io.Writer {
	return &sendWriter{sender: sender}
}

// WriteBufferSize is the largest []byte that Write() will pass to its
// underlying send function. This value can be changed at runtime using
// the GITALY_STREAMIO_WRITE_BUFFER_SIZE environment variable.
var WriteBufferSize = 2

type sendWriter struct {
	sender func([]byte) error
}

func (sw *sendWriter) Write(p []byte) (int, error) {
	var sent int

	for len(p) > 0 {
		chunkSize := len(p)
		if chunkSize > WriteBufferSize {
			chunkSize = WriteBufferSize
		}

		if err := sw.sender(p[:chunkSize]); err != nil {
			return sent, err
		}

		sent += chunkSize
		p = p[chunkSize:]
	}

	return sent, nil
}
