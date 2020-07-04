package linebuffer

import (
	"testing"
)

func TestBuildLineBuffer(t *testing.T) {
	buffer := NewLineBuffer(3)

	if buffer.count != 0 {
		t.Errorf("Unnexpected count: %d", buffer.count)
	}

	if buffer.index != -1 {
		t.Errorf("Unnexpected index: %d", buffer.index)
	}

	if buffer.maxSize != 3 {
		t.Errorf("Unnexpected max size: %d", buffer.maxSize)
	}

	if len(buffer.lines) != 3 {
		t.Errorf("Unnexpected buffer size: %d", len(buffer.lines))
	}
}

func TestPushSingleLine(t *testing.T) {
	buffer := NewLineBuffer(3)

	buffer.Push("line 1")

	if buffer.count != 1 {
		t.Error("Unexpected count")
	}

	if buffer.index != 0 {
		t.Error("Unnexpected index")
	}

	if buffer.lines[0] != "line 1" {
		t.Error("Unnexpected line value")
	}

	line, err := buffer.Get(0)
	if err != nil {
		t.Error(err)
	}

	if line != "line 1" {
		t.Error("Unnexpected line value")
	}
}

func TestPushOverwrites(t *testing.T) {
	buffer := NewLineBuffer(3)

	buffer.Push("line 1")
	buffer.Push("line 2")
	buffer.Push("line 3")

	assertCall := func(offset int, expected string) {
		line, err := buffer.Get(offset)
		if err != nil {
			t.Error(err)
		}

		if line != expected {
			t.Errorf("Expected line: %s, got: %s", expected, line)
		}
	}

	assertCall(0, "line 3")
	assertCall(1, "line 2")
	assertCall(2, "line 1")

	buffer.Push("line 4")

	assertCall(0, "line 4")
	assertCall(1, "line 3")
	assertCall(2, "line 2")
}
