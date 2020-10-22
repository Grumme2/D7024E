package main

import (
	"github.com/Grumme2/D7024E/internal/d7024e"
	//"fmt"
)

func main() {
	me := d7024e.NewContact(d7024e.NewRandomKademliaID(), "localhost")
	rt := d7024e.NewRoutingTable(me)
	network := d7024e.NewNetwork(rt)
	kademlia := d7024e.NewKademlia(&network)
	
	go network.Listen()
	go network.CheckNodesAwaitingResponse()

	cli := d7024e.NewCli(&kademlia)
	cli.AwaitCommand()
}
