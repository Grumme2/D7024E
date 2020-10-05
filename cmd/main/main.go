package main

import (
	"github.com/Grumme2/D7024E/internal/d7024e"
	"fmt"
)

func main() {
	network := d7024e.NewNetwork()
	contactSelf := d7024e.NewContact(d7024e.NewRandomKademliaID(), "127.0.0.1")
	rpcMessage := d7024e.NewRPC(contactSelf, "localhost", "PING", "hello")
	ping := network.SendPingMessage(rpcMessage)
	fmt.Println(ping)
}
