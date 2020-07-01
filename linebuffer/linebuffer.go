package linebuffer

import (
	"fmt"
)

// LineBuffer Ring buffer to hold a number of string "lines" for later recall.
// As lines are pushed to this buffer, the oldest are overwritten as the count reaches max size of the buffer.
// Lines are recalled with Get relative to the last pushed line.
type LineBuffer struct {
	maxSize int
	count   int
	index   int
	lines   []string
}

// NewLineBuffer Construct a new line buffer.
func NewLineBuffer(maxSize int) *LineBuffer {
	buffer := LineBuffer{
		maxSize: maxSize,
		index:   -1,
		count:   0,
		lines:   make([]string, maxSize),
	}

	return &buffer
}

// Count Get the number of lines stored in the buffer.
func (l *LineBuffer) Count() int {
	return l.count
}

// Push Push a line to the buffer.
func (l *LineBuffer) Push(line string) {
	l.lines[l.increaseIdx()] = line
}

// Get Get the line nth last line pushed to the buffer.
func (l *LineBuffer) Get(position int) (string, error) {
	idx, err := l.getIdxForPosition(position)
	if err != nil {
		return "", err
	}

	return l.lines[idx], nil
}

// Increment the head index by 1 making sure to wrap to the beginning.
func (l *LineBuffer) increaseIdx() int {
	l.index++
	if l.index >= l.maxSize {
		l.index = 0
	}

	if l.count < l.maxSize {
		l.count++
	}

	return l.index
}

// Get the idx for `offset` lines ago.
// i.e 0 => last pushed line, 1 => line pushed before that...
func (l *LineBuffer) getIdxForPosition(offset int) (int, error) {
	if offset >= l.count || offset < 0 {
		return -1, fmt.Errorf("Offset %d out of range for size %d", offset, l.count)
	}

	return (l.index - offset + l.maxSize) % l.maxSize, nil
}
