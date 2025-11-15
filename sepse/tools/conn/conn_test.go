package conn

import (
	"bufio"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	// setup
	recv, send := net.Pipe()
	go func() { fmt.Fprint(send, "text\r\n") }()

	// success
	text, err := Read(recv, 1*time.Second)
	assert.Equal(t, "text\r\n", text)
	assert.NoError(t, err)

	// setup
	recv.Close()

	// failure - cannot read
	text, err = Read(recv, 1*time.Second)
	assert.Empty(t, text)
	assert.EqualError(t, err, `cannot read from "pipe" - io: read/write on closed pipe`)
}

func TestWrite(t *testing.T) {
	// setup
	send, recv := net.Pipe()

	// success
	go func() {
		err := Write(send, 1*time.Second, "text\r\n")
		assert.NoError(t, err)
	}()

	// confirm - received
	text, err := bufio.NewReader(recv).ReadString('\n')
	assert.Equal(t, "text\r\n", text)
	assert.NoError(t, err)

	// setup
	send.Close()

	// failure - cannot write
	err = Write(send, 1*time.Second, "text\r\n")
	assert.EqualError(t, err, `cannot write into "pipe" - io: read/write on closed pipe`)
}
