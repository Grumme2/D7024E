package main

import (
	"github.com/Grumme2/D7024E/internal/d7024e"
)

func main() {
	network := d7024e.NewNetwork()
	contactSelf := d7024e.NewContact(d7024e.NewRandomKademliaID(), "127.0.0.1")
	network.Listen(contactSelf)
}
