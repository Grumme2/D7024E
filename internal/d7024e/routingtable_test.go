package d7024e

import (
	"testing"

	"github.com/stretchr/testify/assert"
	//	"strconv"
)

func TestRoutingTable(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"
	str3 := "1111111200000000000000000000000000000000"
	str4 := "1111111300000000000000000000000000000000"
	str5 := "1111111400000000000000000000000000000000"
	str6 := "2111111400000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	id3 := NewKademliaID(&str3)
	id4 := NewKademliaID(&str4)
	id5 := NewKademliaID(&str5)
	id6 := NewKademliaID(&str6)

	rt := NewRoutingTable(NewContact(&meid, "localhost:8001"))
	con1 := NewContact(&id1, "localhost:8001")
	con2 := NewContact(&id2, "localhost:8002")
	con3 := NewContact(&id3, "localhost:8002")
	con4 := NewContact(&id4, "localhost:8002")
	con5 := NewContact(&id5, "localhost:8002")
	con6 := NewContact(&id6, "localhost:8001")

	rt.AddContact(con1)
	rt.AddContact(con2)
	rt.AddContact(con3)
	rt.AddContact(con4)
	rt.AddContact(con5)
	rt.AddContact(con6)

	contacts := rt.FindClosestContacts(&id6, 20)

	// fmt.Println(contacts[0].String())
	// fmt.Println(contacts[1].String())
	// fmt.Println(contacts[2].String())
	// fmt.Println(contacts[3].String())
	// fmt.Println(contacts[4].String())
	// fmt.Println(contacts[5].String())

	assert.Equal(t, contacts[0].String(), `contact("b4bfb0ce7e0e6f16c2acc5facde2ae723210b66a", "localhost:8001")`)
	assert.Equal(t, contacts[1].String(), `contact("b433805b2c8d792c75e5b069f217dfcbc0f4517a", "localhost:8001")`)
	assert.Equal(t, contacts[2].String(), `contact("b1221d78ddcecd0981b8847572aa2f2cdd09fcc0", "localhost:8002")`)
	assert.Equal(t, contacts[3].String(), `contact("bb98710a49b8c2a9b90afec87852109fef6adf81", "localhost:8002")`)
	assert.Equal(t, contacts[4].String(), `contact("ef28e8376adc64a62ca0c8e58298cdc4f98b57f6", "localhost:8002")`)
	assert.Equal(t, contacts[5].String(), `contact("53a1cec570516890d2f1ab5c64587d6f773c5548", "localhost:8002")`)

}

func TestRemoveContact(t *testing.T) {
	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"
	str3 := "1111111200000000000000000000000000000000"
	str4 := "1111111300000000000000000000000000000000"
	str5 := "1111111400000000000000000000000000000000"
	str6 := "2111111400000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	id3 := NewKademliaID(&str3)
	id4 := NewKademliaID(&str4)
	id5 := NewKademliaID(&str5)
	id6 := NewKademliaID(&str6)

	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con2 := NewContact(&id2, "localhost")
	con3 := NewContact(&id3, "localhost")
	con4 := NewContact(&id4, "localhost")
	con5 := NewContact(&id5, "localhost")
	con6 := NewContact(&id6, "localhost")

	rt.AddContact(con1)
	rt.AddContact(con2)
	rt.AddContact(con3)
	rt.AddContact(con4)
	rt.AddContact(con5)
	rt.AddContact(con6)
	index := rt.GetBucketIndex(&id3)
	rt.buckets[index].RemoveContact(con3)
	conInBucket := rt.buckets[index].IsContactInBucket(con3)
	if conInBucket {
		t.Errorf("Cant remove contact!")
	}
}

func TestFindClosestContacts(t *testing.T) {

	mestr := "FFFFFFFF00000000000000000000000000000000"
	str1 := "FFFFFFFF00000000000000000000000000000000"
	str2 := "1111111100000000000000000000000000000000"
	str3 := "1111111200000000000000000000000000000000"
	str4 := "2111111300000000000000000000000000000000"
	str6 := "2111111400000000000000000000000000000000"

	meid := NewKademliaID(&mestr)
	id1 := NewKademliaID(&str1)
	id2 := NewKademliaID(&str2)
	id3 := NewKademliaID(&str3)
	id4 := NewKademliaID(&str4)
	id6 := NewKademliaID(&str6)

	rt := NewRoutingTable(NewContact(&meid, "localhost"))
	con1 := NewContact(&id1, "localhost")
	con := NewContact(&id2, "localhost")
	notEvenClose := NewContact(&id3, "localhost")
	nextClosest := NewContact(&id4, "localhost")
	closestCon := NewContact(&id6, "localhost")

	rt.AddContact(con1)
	rt.AddContact(con)
	rt.AddContact(notEvenClose)
	rt.AddContact(nextClosest)
	rt.AddContact(closestCon)
	// test := "closestCon: " + closestCon.String()
	// fmt.Println(test)
	// test2 := "nextclosest: " + nextClosest.String()
	// fmt.Println(test2)
	// test3 := "notevenclose: " + notEvenClose.String()
	// fmt.Println(test3)
	// test4 := "con: " + con.String()
	// fmt.Println(test4)

	close := rt.FindClosestContacts(&id6, 2)
	assert.Equal(t, close[0].ID, closestCon.ID)
	assert.Equal(t, close[1].ID, notEvenClose.ID)
	rt.RemoveContact(closestCon)
	newClose := rt.FindClosestContacts(&id6, 2)
	assert.Equal(t, newClose[0].ID, notEvenClose.ID)
	assert.Equal(t, newClose[1].ID, con.ID)
}
