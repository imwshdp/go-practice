package solution

import (
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	proc := MockProcess{}
	go proc.Run()
	gracefulShutdown(&proc)
}

func gracefulShutdown(
	proc *MockProcess,
) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT)
	<-sigChan

	done := make(chan struct{})
	defer close(done)

	go func(proc *MockProcess, done chan struct{}) {
		proc.Stop()
		done <- struct{}{}
	}(proc, done)

	select {
	case <-done: // proc.Stop() finished
	case <-time.After(5 * time.Second): // timeout shutdown
	case <-sigChan: // SIGINT forced shutdown
		return
	}
}
