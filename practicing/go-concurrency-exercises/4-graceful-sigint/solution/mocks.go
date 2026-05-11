package solution

import "sync"

type MockProcess struct {
	mu        sync.Mutex
	isRunning bool
}

func (m *MockProcess) Run() {}

func (m *MockProcess) Stop() {}
