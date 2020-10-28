package main

import (
	"fmt"
	"strings"

	"github.com/Grumme2/D7024E/internal/d7024e"
	//"fmt"
)

func main() {
	meid := d7024e.NewRandomKademliaID()
	ip := d7024e.GetLocalIP()
	splitIP := strings.Split(ip, ".")
	fmt.Println(splitIP)
<<<<<<< HEAD
	if splitIP[3] == "3" {
		mestr := "2111111300000000000000000000123000000000"
		meid2 := d7024e.NewKademliaID(&mestr)
		meid = &meid2
=======
	if splitIP[3] == "3" { // if ip ends with a 3 its the bootstrapnode
		mestr := "2111111300000000000000000000123000000000" // hard coded bootstrap node id
		meid = d7024e.NewKademliaID(&mestr)
>>>>>>> some changes
	}
	me := d7024e.NewContact(meid, ip)
	rt := d7024e.NewRoutingTable(me)
	network := d7024e.NewNetwork(rt)
	kademlia := d7024e.NewKademlia(&network)
<<<<<<< HEAD

	go network.Listen()
	go network.CheckNodesAwaitingResponse()
	// kademlia.JoinNetwork()
=======
	kademlia.JoinNetwork()
>>>>>>> some changes
	fmt.Println(me.Address)
	// go network.Listen()
	// go network.CheckNodesAwaitingResponse()

	cli := d7024e.NewCli(&kademlia)
	cli.AwaitCommand()
}
