package d7024e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	network := NewNetwork(rt)
	assert.NotNil(t, network)

}

func TestGetLocalIP(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	network := NewNetwork(rt)
	assert.Equal(t, network.GetLocalIP(), "130.240.64.55")

}

func TestKriplesJSON(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	network := NewNetwork(rt)
	con1 := NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000001"), "localhost:8003")
	con2 := NewContact(NewKademliaID("1111111100000000000000000000000000000001"), "localhost:8004")
	con3 := NewContact(NewKademliaID("1111111200000000000000000000000000000001"), "localhost:8006")
	list := []Contact{con1, con2, con3}
	//fmt.Println(network.KTriplesJSON(list))
	assert.Equal(t, network.KTriplesJSON(list), `[{"ID":[255,255,255,255,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1],"Address":"localhost:8003","KeyValueStore":{}},{"ID":[17,17,17,17,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1],"Address":"localhost:8004","KeyValueStore":{}},{"ID":[17,17,17,18,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1],"Address":"localhost:8005","KeyValueStore":{}}]`)
}

// func TestSendMessage(t *testing.T) {
// rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))
// rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
// con1 := NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
// rt.AddContact(con1)
// network := NewNetwork(rt)
// pingRPC := NewRPC(network.routingTable.me, con1.Address, "PING", "")
// assert.Equal(t, network.SendMessage(pingRPC), true)
//
//
// }
