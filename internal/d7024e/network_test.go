package d7024e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNetwork(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)

	network := NewNetwork(rt)
	assert.NotNil(t, network)

}

// func TestGetLocalIP(t *testing.T) {
// mestr := "FFFFFFFF00000000000000000000000000000000"
// str1 := "FFFFFFFF00000000000000000000000000000000"
// str2 := "1111111100000000000000000000000000000000"
//
// meid := NewKademliaID(&mestr)
// id1 := NewKademliaID(&str1)
// id2 := NewKademliaID(&str2)
// rt := NewRoutingTable(NewContact(&meid, "localhost"))
// con1 := NewContact(&id1, "localhost")
// con2 := NewContact(&id2, "localhost")
// network := NewNetwork(rt)
// assert.Equal(t, network.GetLocalIP(), "192.168.1.12")
//
// }

func TestKriplesJSON(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)
	network := NewNetwork(rt)

	str3 := "FFFFFFFF00000000000000000000000000000001"
	str4 := "1111111100000000000000000000000000000001"
	str5 := "1111111200000000000000000000000000000001"

	id3 := NewKademliaID(&str3)
	id4 := NewKademliaID(&str4)
	id5 := NewKademliaID(&str5)

	con3 := NewContact(&id3, "localhost")
	con4 := NewContact(&id4, "localhost")
	con5 := NewContact(&id5, "localhost")

	list := []Contact{con3, con4, con5}
	//fmt.Println(network.KTriplesJSON(list))
	assert.Equal(t, network.KTriplesJSON(list), `[{"ID":[225,19,161,37,135,218,181,90,72,245,226,134,60,209,249,0,53,45,185,171],"Address":"localhost","KeyValueStore":{}},{"ID":[95,50,45,254,97,219,19,205,45,221,20,19,145,116,84,112,176,160,214,107],"Address":"localhost","KeyValueStore":{}},{"ID":[101,139,88,248,141,101,156,47,114,72,75,13,78,108,251,16,113,172,122,224],"Address":"localhost","KeyValueStore":{}}]`)
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

func TestJSONLookUpDataResponse(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)
	network := NewNetwork(rt)

	str3 := "FFFFFFFF00000000000000000000000000000001"
	id3 := NewKademliaID(&str3)
	con3 := NewContact(&id3, "localhost")

	lookUpDataResponse := LookUpDataResponse{true, "test", con3}

	str := network.JSONEncodeLookUpDataResponse(lookUpDataResponse)
	fin := network.JSONDecodeLookUpDataResponse(str)

	assert.Equal(t, fin.DataFound, true)
	assert.Equal(t, fin.Data, "test")
	assert.Equal(t, fin.Node, con3)

}

func TestSendXMessage(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)
	network := NewNetwork(rt)

	str3 := "FFFFFFFF00000000000000000000000000000001"
	id3 := NewKademliaID(&str3)
	con3 := NewContact(&id3, "localhost")

	lookUpDataResponse := LookUpDataResponse{true, "test", con3}
	network.lookUpDataResponse = lookUpDataResponse

	databool, datastr, datacon := network.SendFindDataMessage()

	assert.Equal(t, databool, true)
	assert.Equal(t, datastr, "test")
	assert.Equal(t, datacon, con3)

	network.lookUpContactResponse = LookUpContactResponse{"test"}
	constr := network.SendFindContactMessage()

	assert.Equal(t, constr, "test")

	network.pingResponse = PINGResponse{true, "test"}

	pingbool, pingstr := network.SendPINGMessage()

	assert.Equal(t, pingbool, true)
	assert.Equal(t, pingstr, "test")

}

func TestGetLocalIP(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)
	// network := NewNetwork(rt)

	assert.Equal(t, GetLocalIP(), "130.240.65.43")
}

func TestAddLookforData(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	rt.AddContact(con1)
	rt.AddContact(con2)
	network := NewNetwork(rt)

	hash := network.AddToStore("test")
	databool, element := network.LookForData(hash)
	assert.True(t, databool)
	assert.Equal(t, "test", element)
}
