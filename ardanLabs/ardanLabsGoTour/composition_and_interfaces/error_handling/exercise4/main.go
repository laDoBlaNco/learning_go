package example4

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

/*
	The problem we noted in the last example is that we are no longer decoupled from the
	concrete error value. This means if the concrete error value is changed, then our code
	can break. The beautiful part of using an interface here for error  handling is being
	decoupled from breaking changes.

	If the concrete error value has a method set, then you can use an interface for the
	type check. As an example, the net package has many concrete error types that implement
	different methods. One common method is called Temporary. this method allows the user
	to test if the networking error is critical or just something that can recover on its
	own
*/

// client reps a single connection in the room
type client struct {
	name   string
	reader *bufio.Reader
}

// TypeAsContext shows how to check multiple types of possible custom error types
// that can be returned from the net package
func (c *client) TypeAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case *net.OpError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}
			case *net.AddrError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}
			case *net.DNSConfigError:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}
				log.Println("read-routine", err)
			}
		}
		fmt.Println(line)
	}
}

// temporary is declared to test for the existence of the method coming from the net package
type temporary interface {
	Temporary() bool
}

// In this code, the call to ReadString could fail with an error from the net package. In this
// case, an interface is declared that reps the common behavior a given concrete error value
// could implement. Then with a generic type assertion, you test if that behavior exists.
// and you can call into it. The best part, you stay in a decoupled state with our error
// handling.
func (c *client) BehaviorAsContext() {
	for {
		line, err := c.reader.ReadString('\n')
		if err != nil {
			switch e := err.(type) {
			case temporary:
				if !e.Temporary() {
					log.Println("Temporary: Client leaving chat")
					return
				}
			default:
				if err == io.EOF {
					log.Println("EOF: Client leaving chat")
					return
				}
				log.Println("read-routine", err)
			}
		}
		fmt.Println(line)
	}
}
