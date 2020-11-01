package d7024e

import (
	"fmt"
	"testing"
	"time"
	//"reflect"
	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	id1 := NewRandomKademliaID()
	id2 := NewRandomKademliaID()
	id3 := NewRandomKademliaID()

	con1 := NewContact(&id1, "localhost:8000")
	con2 := NewContact(&id2, "localhost:8001")
	con3 := NewContact(&id3, "localhost:8002")

	lst := []Contact{con1, con2, con3}
	check := in(con3, lst)
	if !check {
		t.Errorf("CONTACT NOT IN LIST")
	}
	assert.True(t, check)
}

func TestNodeLookup(t *testing.T) {
	sender := NewRandomKademliaID()
    time.Sleep(2 * time.Millisecond)
    recivier := NewRandomKademliaID()
    time.Sleep(2 * time.Millisecond)
	rtContactId := NewRandomKademliaID()
	
	me := NewContact(&sender, "localhost:8000")
	rtContact := NewContact(&rtContactId, "localhost:8090")
	rt := NewRoutingTable(me)
	network := NewNetwork(rt)
	kademlia := NewKademlia(&network)
	kademlia.network.routingTable.AddContact(NewContact(&recivier, "localhost:8001"))
	go network.ListenHandler()
	go kademlia.LookupContact(&rtContact)
	msg := <- kademlia.network.external
	fmt.Printf(string(msg))
}

// func TestLookupData(t *testing.T) {
// rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost"))
// con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost")
// con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost")
// con3 := NewContact(NewKademliaID("1111111100000000000200000000000000000000"), "localhost")
// con4 := NewContact(NewKademliaID("1111111100000000000300000000000000000000"), "localhost")
// con5 := NewContact(NewKademliaID("1111111100000000000400000000000000000000"), "localhost")
// con6 := NewContact(NewKademliaID("1111111100000000000500000000000000000000"), "localhost")
// con7 := NewContact(NewKademliaID("1111111100000000000600000000000000000000"), "localhost")
// rt.AddContact(con1)
// rt.AddContact(con2)
// rt.AddContact(con3)
// rt.AddContact(con4)
// rt.AddContact(con5)
// rt.AddContact(con6)
// rt.AddContact(con7)
// network := NewNetwork(rt)
// kademlia := Kademlia{network}
// shortlist := kademlia.LookupData(&con2)
// fmt.Println(shortlist)
// assert.Equal(t, con2.ID, shortlist[0].ID)
//
// }
//
