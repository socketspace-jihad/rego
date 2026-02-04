package server

import "testing"

type MockServer struct{}

func (m *MockServer) Serve() {}

func NewMockServer() Server {
	return &MockServer{}
}

func Test_Server(t *testing.T) {
	t.Run("register a server", func(t *testing.T) {
		RegisterServer("mock", NewMockServer())
	})

	t.Run("get a server", func(t *testing.T) {
		srv, err := GetServer("mock")
		if err != nil {
			t.Error(err)
		}
		srv.Serve()
	})
}
