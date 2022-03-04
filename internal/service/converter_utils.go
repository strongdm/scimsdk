package service

import (
	"io"
	"strings"
)

func convertResponseBodyToBytes(body io.ReadCloser) ([]byte, error) {
	buff := new(strings.Builder)
	_, err := io.Copy(buff, body)
	if err != nil {
		return nil, err
	}
	return []byte(buff.String()), nil
}
