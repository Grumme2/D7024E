package d7024e

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	con := NewContact(NewRandomKademliaID(), "localhost:8000")
	con2 := NewContact(NewRandomKademliaID(), "localhost:8001")
	con3 := NewContact(NewRandomKademliaID(), "localhost:8002")
	lst := []Contact{con, con2, con3}
	check := in(con3, lst)
	if !check {
		t.Errorf("CONTACT NOT IN LIST")
	}
	assert.True(t, check)
}

func TestLookupContact(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"
	str3 := "1111111100000000000200000000000000000000"
	str4 := "1111111100000000000300000000000000000000"
	str5 := "1111111100000000000400000000000000000000"
	str6 := "1111111100000000000500000000000000000000"
	str7 := "1111111100000000000600000000000000000000"
	str8 := "1111111100000000000700000000000000000000"
	str9 := "1111111100000000000800000000000000000000"
	str10 := "1111111100000000002400000000000000000000"
	str11 := "1111111100000000000900000000000000000000"
	str12 := "1111111100000000001000000000000000000000"
	str13 := "1111111100000000001100000000000000000000"
	str14 := "1111111100000000001200000000000000000000"
	str15 := "1111111100000000001300000000000000000000"
	str16 := "1111111100000000001400000000000000000000"
	str17 := "1111111100000000001500000000000000000000"
	str18 := "11111111000000000016100000000000000000000"
	str19 := "1111111100000000001700000000000000000000"
	str20 := "1111111100000000001800000000000000000000"
	str21 := "1111111100000000001900000000000000000000"
	str22 := "1111111100000000002000000000000000000000"
	str23 := "1111111100000000002100000000000000000000"
	str24 := "1111111100000000002200000000000000000000"
	str25 := "1111111100000000002300000000000000000000"
	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	id3 := NewKademliaID(&str3)
	id4 := NewKademliaID(&str4)
	id5 := NewKademliaID(&str5)
	id6 := NewKademliaID(&str6)
	id7 := NewKademliaID(&str7)
	id8 := NewKademliaID(&str8)
	id9 := NewKademliaID(&str9)
	id10 := NewKademliaID(&str10)
	id11 := NewKademliaID(&str11)
	id12 := NewKademliaID(&str12)
	id13 := NewKademliaID(&str13)
	id14 := NewKademliaID(&str14)
	id15 := NewKademliaID(&str15)
	id16 := NewKademliaID(&str16)
	id17 := NewKademliaID(&str17)
	id18 := NewKademliaID(&str18)
	id19 := NewKademliaID(&str19)
	id20 := NewKademliaID(&str20)
	id21 := NewKademliaID(&str21)
	id22 := NewKademliaID(&str22)
	id23 := NewKademliaID(&str23)
	id24 := NewKademliaID(&str24)
	id25 := NewKademliaID(&str25)
	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	con3 := NewContact(&id3, "localhost")
	con4 := NewContact(&id4, "localhost")
	con5 := NewContact(&id5, "localhost")
	con6 := NewContact(&id6, "localhost")
	con7 := NewContact(&id7, "localhost")
	con8 := NewContact(&id8, "localhost")
	con9 := NewContact(&id9, "localhost")
	con10 := NewContact(&id10, "localhost")
	con11 := NewContact(&id11, "localhost")
	con12 := NewContact(&id12, "localhost")
	con13 := NewContact(&id13, "localhost")
	con14 := NewContact(&id14, "localhost")
	con15 := NewContact(&id15, "localhost")
	con16 := NewContact(&id16, "localhost")
	con17 := NewContact(&id17, "localhost")
	con18 := NewContact(&id18, "localhost")
	con19 := NewContact(&id19, "localhost")
	con20 := NewContact(&id20, "localhost")
	con21 := NewContact(&id21, "localhost")
	con22 := NewContact(&id22, "localhost")
	con23 := NewContact(&id23, "localhost")
	con24 := NewContact(&id24, "localhost")
	con25 := NewContact(&id25, "localhost")

	rt.AddContact(con1)
	rt.AddContact(con2)
	rt.AddContact(con3)
	rt.AddContact(con4)
	rt.AddContact(con5)
	rt.AddContact(con6)
	rt.AddContact(con7)
	rt.AddContact(con8)
	rt.AddContact(con9)
	rt.AddContact(con10)
	rt.AddContact(con11)
	rt.AddContact(con12)
	rt.AddContact(con13)
	rt.AddContact(con14)
	rt.AddContact(con15)
	rt.AddContact(con16)
	rt.AddContact(con17)
	rt.AddContact(con18)
	rt.AddContact(con19)
	rt.AddContact(con20)
	rt.AddContact(con21)
	rt.AddContact(con22)
	rt.AddContact(con23)
	rt.AddContact(con24)
	rt.AddContact(con25)
	network := NewNetwork(rt)
	go network.Listen()
	kademlia := Kademlia{&network}
	// shortlist := kademlia.LookupContact(&con17)
	// assert.Equal(t, con17.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con18)
	// assert.Equal(t, con18.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con19)
	// assert.Equal(t, con19.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con20)
	// assert.Equal(t, con20.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con21)
	// assert.Equal(t, con21.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con22)
	// assert.Equal(t, con22.ID, shortlist[0].ID)
	shortlistJSON := kademlia.LookupContact(&con23)
	shortlist := network.KTriples(shortlistJSON)
	//fmt.Println(shortlist)
	assert.Equal(t, con23.ID, shortlist[0].ID)
	// //shortlist = kademlia.LookupContact(&con24)
	// //assert.Equal(t, con24.ID, shortlist[0].ID)
	// //shortlist = kademlia.LootkupContact(&con25)
	// //assert.Equal(t, con25.ID, shortlist[0].ID)
	// //shortlist = kademlia.LookupContact(&con26)
	// //assert.Equal(t, con26.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con1)
	// assert.Equal(t, con1.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con2)
	// assert.Equal(t, con2.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con3)
	// assert.Equal(t, con3.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con4)
	// assert.Equal(t, con4.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con5)
	// assert.Equal(t, con5.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con6)
	// assert.Equal(t, con6.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con7)
	// assert.Equal(t, con7.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con8)
	// assert.Equal(t, con8.ID, shortlist[0].ID)
	// shortlist = kademlia.LookupContact(&con10)
	// assert.Equal(t, con10.ID, shortlist[0].ID)

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
