package conveyer

func New[T comparable](in <-chan T, fn func(T) T) <-chan T {
	out := make(chan T)

	go func() {
		defer close(out)

		for {
			val, ok := <-in
			if !ok {
				return
			}
			out <- fn(val)
		}
	}()

	return out
}
