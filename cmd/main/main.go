package main

import (

	"github.com/Grumme2/D7024E/internal/d7024e"
	//"fmt"
)

func main() {
	contactSelf := d7024e.NewContact(d7024e.NewRandomKademliaID(), "127.0.0.1")
	rt := d7024e.NewRoutingTable(contactSelf)
	network := d7024e.NewNetwork(rt)
	network.TestCreateAwaitingReponseObjects()
	network.CheckNodesAwaitingResponse()

	/*fmt.Println(time.Now().Unix())
	network := d7024e.NewNetwork()
	contactSelf := d7024e.NewContact(d7024e.NewRandomKademliaID(), "127.0.0.1")
	rt := d7024e.NewRoutingTable(contactSelf)
	rpcMessage := d7024e.NewRPC(contactSelf, "localhost", "PING", "hello")
	network.Listen(*rt)
	ping := network.SendPingMessage(rpcMessage)
	fmt.Println(ping)*/

}
