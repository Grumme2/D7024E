package d7024e

import (
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
//	"strconv"
)

func TestRoutingTable(t *testing.T) {
	rt := NewRoutingTable(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8000"))

	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111300000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("1111111400000000000000000000000000000000"), "localhost:8002"))
	rt.AddContact(NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002"))

	contacts := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 20)
	for i := range contacts {
		fmt.Println(contacts[i].String())
	}
}

func TestRemoveContact(t *testing.T) {
	kadID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	me := NewContact(kadID, "localhost:8000")
	rt := NewRoutingTable(me)
	con:= NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(con)
	rt.AddContact(NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002"))
	index := rt.GetBucketIndex(kadID)
	rt.buckets[index].RemoveContact(con)
	conInBucket:=rt.buckets[index].IsContactInBucket(con)
	if conInBucket {
		t.Errorf("Cant remove contact!")
	}
}

func TestFindClosestContacts(t *testing.T) {
	kadID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	me := NewContact(kadID, "localhost:8000")
	rt := NewRoutingTable(me)
	con:= NewContact(NewKademliaID("1111111100000000000000000000000000000000"), "localhost:8002")
	rt.AddContact(NewContact(NewKademliaID("FFFFFFFF00000000000000000000000000000000"), "localhost:8001"))
	rt.AddContact(con)
	closestCon :=NewContact(NewKademliaID("2111111400000000000000000000000000000000"), "localhost:8002")
	nextClosest :=NewContact(NewKademliaID("2111111300000000000000000000000000000000"), "localhost:8002")
	notEvenClose := NewContact(NewKademliaID("1111111200000000000000000000000000000000"), "localhost:8002")
	rt.AddContact(notEvenClose)
	rt.AddContact(closestCon)
	rt.AddContact(nextClosest)
	close :=rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 2)
	assert.Equal(t,close[0].ID ,closestCon.ID)
	assert.Equal(t,close[1].ID, nextClosest.ID)
	rt.RemoveContact(closestCon)
	newClose := rt.FindClosestContacts(NewKademliaID("2111111400000000000000000000000000000000"), 2)
	assert.Equal(t, newClose[0].ID, nextClosest.ID)
	assert.Equal(t, newClose[1].ID, con.ID)
}
