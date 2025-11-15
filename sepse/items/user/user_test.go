package user

import (
	"bufio"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// setup
	conn, _ := net.Pipe()

	// success
	user := New(conn, "LANE")
	assert.Equal(t, conn, user.Conn)
	assert.Equal(t, map[string]bool{"LANE": true}, user.Lanes)
}

func TestClose(t *testing.T) {
	// setup
	conn, _ := net.Pipe()
	user := New(conn, "LANE")

	// success
	err := user.Close()
	assert.NoError(t, err)

	// confirm - conn closed
	_, err = user.Conn.Write(nil)
	assert.Error(t, err)
}

func TestRead(t *testing.T) {
	// setup
	conn, pipe := net.Pipe()
	user := New(conn, "LANE")
	go func() { pipe.Write([]byte("text\n")) }()

	// success
	text, err := user.Read()
	assert.Equal(t, "text\n", text)
	assert.NoError(t, err)
}

func TestSubscribe(t *testing.T) {
	// setup
	user := New(nil)

	// success
	user.Subscribe("LANE")
	assert.Equal(t, map[string]bool{"LANE": true}, user.Lanes)
}

func TestUnsubscribe(t *testing.T) {
	// setup
	user := New(nil, "LANE")

	// success
	user.Unsubscribe("LANE")
	assert.Empty(t, user.Lanes)
}

func TestWrite(t *testing.T) {
	// setup
	conn, pipe := net.Pipe()
	user := New(conn, "LANE")

	// success
	go func() {
		err := user.Write("%s\n", "text")
		assert.NoError(t, err)
	}()

	// confirm - text received
	text, err := bufio.NewReader(pipe).ReadString('\n')
	assert.Equal(t, "text\n", text)
	assert.NoError(t, err)
}

func TestWriteLane(t *testing.T) {
	// setup
	conn, pipe := net.Pipe()
	pipe.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	user := New(conn, "LANE")

	// success - lane exists
	go func() {
		err := user.WriteLane("LANE", "%s\n", "text")
		assert.NoError(t, err)
	}()

	// confirm - text received
	text, err := bufio.NewReader(pipe).ReadString('\n')
	assert.Equal(t, "text\n", text)
	assert.NoError(t, err)

	// success - lane does not exist
	err = user.WriteLane("NOPE", "%s\n", "text")
	assert.NoError(t, err)

	// confirm - text not received
	_, err = pipe.Read(nil)
	assert.Error(t, err)
}
