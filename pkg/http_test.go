package pkg

import (
	"errors"
	"testing"
)

func Test_HTTPRegoSDK(t *testing.T) {
	t.Run("create http rego connection", func(t *testing.T) {
		conn := NewHTTPRegoConnection(nil)
		if conn == nil {
			t.Error(errors.New("http rego doesn't initialize correctly"))
			return
		}

		if conn.Transport == nil {
			t.Error(errors.New("transport value is nil"))
			return
		}

		if conn.Client == nil {
			t.Error(errors.New("http client is nil"))
			return
		}
	})
}
