package main

import (
	"fmt"
	"strings"

	"github.com/Grumme2/D7024E/internal/d7024e"
)

func main() {
	meid := d7024e.NewRandomKademliaID()
	ip := d7024e.GetLocalIP()
	splitIP := strings.Split(ip, ".")
	fmt.Println(splitIP)
	if splitIP[3] == "3" {
		mestr := "2111111300000000000000000000123000000000"
		meid = d7024e.NewKademliaID(&mestr)
	}
	me := d7024e.NewContact(&meid, ip)
	rt := d7024e.NewRoutingTable(me)
	network := d7024e.NewNetwork(rt)
	kademlia := d7024e.NewKademlia(&network)

	go network.Listen()
	go network.CheckNodesAwaitingResponse()
	// kademlia.JoinNetwork()

	fmt.Println(me.Address)
	// go network.Listen()
	// go network.CheckNodesAwaitingResponse()

	cli := d7024e.NewCli(&kademlia)
	cli.AwaitCommand()
}
