package solution

import "time"

const timeLimit = 10 * time.Second

type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64
}

func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	done := make(chan struct{})

	go func(done chan struct{}) {
		process()
		close(done)
	}(done)

	select {
	case <-time.After(timeLimit):
		return false
	case <-done:
		return true
	}
}
