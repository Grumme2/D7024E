package main

import (
	"fmt"

	"github.com/Grumme2/D7024E/internal/d7024e"
	"github.com/Grumme2/D7024E/internal/test"
)

func main() {
	fmt.Println("Hello, World!")
	test.TestPrint()
	rt := d7024e.NewRoutingTable(d7024e.NewContact(d7024e.NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
	contact := d7024e.NewContact(d7024e.NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8001")
	rt.AddContact(contact)
	ka := d7024e.Kademlia{}
	ka.Rt = *rt
	ka.LookupContact(&contact)
}
