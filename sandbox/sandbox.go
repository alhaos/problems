package sandbox

import (
	"bytes"
	"errors"
	"io"
)

func Contains(r io.Reader, seq []byte) (bool, error) {

	l := len(seq)

	slidingBuffer := make([]byte, l*2)
	readBuffer := make([]byte, l)

	for {
		n, err := r.Read(readBuffer)
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return false, err
		}

		slidingBuffer = append(slidingBuffer, readBuffer...)
		slidingBuffer = slidingBuffer[min(l, n):]

		if bytes.Contains(slidingBuffer, seq) {
			return true, nil
		}
	}

	return false, nil
}
