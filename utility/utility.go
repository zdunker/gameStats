package utility

import (
	"bytes"
	"io"
)

func ReadBytes(fp io.Reader) (bs []byte, err error) {
	var buf bytes.Buffer
	_, err = io.Copy(&buf, fp)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), err
}
