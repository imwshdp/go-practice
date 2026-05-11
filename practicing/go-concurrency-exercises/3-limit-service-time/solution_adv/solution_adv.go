package solution_adv

import (
	"sync"
	"time"
)

const timeLimit = 10 * time.Second

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
	mu        sync.Mutex
}

func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		start := time.Now()
		process()

		spendTime := time.Since(start).Seconds()

		u.mu.Lock()
		u.TimeUsed += int64(spendTime)
		u.mu.Unlock()

		return true
	}

	u.mu.Lock()
	if u.TimeUsed >= int64(timeLimit.Seconds()) {
		u.mu.Unlock()
		return false
	}
	u.mu.Unlock()

	done := make(chan struct{})
	start := time.Now()

	go func(done chan struct{}) {
		process()
		close(done)
	}(done)

	u.mu.Lock()
	userTimeLeft := timeLimit - time.Duration(u.TimeUsed)*time.Second
	u.mu.Unlock()

	select {
	case <-time.After(userTimeLeft):
		spendTime := time.Since(start).Seconds()

		u.mu.Lock()
		u.TimeUsed += int64(spendTime)
		u.mu.Unlock()

		return false
	case <-done:
		spendTime := time.Since(start).Seconds()

		u.mu.Lock()
		u.TimeUsed += int64(spendTime)
		u.mu.Unlock()

		return true
	}
}
