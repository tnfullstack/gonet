package listener

import (
	"fmt"
	"net"
	"testing"
)

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp6", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	// defer listener.Close()

	defer func() {
		_ = listener.Close()
	}()

	t.Logf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
		}

		go func(c net.Conn) {
			defer c.Close()
		}(conn)
	}
}
