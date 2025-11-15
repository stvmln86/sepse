package lane

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stvmln86/sepse/sepse/items/note"
)

var (
	mockTime = time.Unix(1234567890, 0).Local()
)

func TestNew(t *testing.T) {
	// setup
	conn, _ := net.Pipe()

	// success
	lane := New("NAME", conn)
	assert.Equal(t, "NAME", lane.Name)
	assert.Contains(t, lane.Conns, conn)
	assert.True(t, lane.Mutex.TryLock())
}

func TestAdd(t *testing.T) {
	// setup
	conn, _ := net.Pipe()
	lane := New("NAME")

	// success
	lane.Add(conn)
	assert.True(t, lane.Conns[conn])
}

func TestClose(t *testing.T) {
	// setup
	conn, _ := net.Pipe()
	lane := New("NAME", conn)

	// success
	err := lane.Close()
	assert.NoError(t, err)

	// confirm - conn closed
	_, err = fmt.Fprintf(conn, "test\n")
	assert.Error(t, err)
}

func TestRemove(t *testing.T) {
	// setup
	conn, _ := net.Pipe()
	lane := New("NAME", conn)

	// success
	lane.Remove(conn)
	assert.Empty(t, lane.Conns)
}

func TestSend(t *testing.T) {
	// setup
	send, recv := net.Pipe()
	lane := New("NAME", send)
	rder := bufio.NewReader(recv)
	note := note.New("body", mockTime)

	// success
	go func() {
		err := lane.Send(note)
		assert.NoError(t, err)
	}()

	// confirm - note received
	text, err := rder.ReadString('\n')
	assert.Equal(t, "NAME 1234567890 body\n", text)
	assert.NoError(t, err)
}
