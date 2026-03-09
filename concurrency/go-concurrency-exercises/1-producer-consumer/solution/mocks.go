package solution

import "errors"

var ErrEOF = errors.New("End of File")

type Tweet struct {
	Username string
	Text     string
}

func (t *Tweet) IsTalkingAboutGo() bool {
	return false
}

type Stream struct {
	pos    int
	tweets []Tweet
}

func (s *Stream) Next() (*Tweet, error) {
	return nil, nil
}

func GetMockStream() Stream {
	return Stream{}
}
