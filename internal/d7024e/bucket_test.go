package d7024e

import (
	"testing"
)

func TestAddToBucket(t *testing.T) {
	buck := NewBucket()
	id1 := NewRandomKademliaID()
	id2 := NewRandomKademliaID()
	id3 := NewRandomKademliaID()
	id4 := NewRandomKademliaID()
	id5 := NewRandomKademliaID()

	conOne := NewContact(&id1, "localhost:8000")
	conTwo := NewContact(&id2, "localhost:8001")
	conThree := NewContact(&id3, "localhost:8002")
	conFour := NewContact(&id4, "localhost:8003")
	conFive := NewContact(&id5, "localhost:8004")

	buck.AddContact(conOne)
	buck.AddContact(conTwo)
	buck.AddContact(conThree)
	buck.AddContact(conFour)
	buck.AddContact(conFive)
	if buck.Len() > bucketSize {
		t.Errorf("Bucket can be larger than bucketSize")
	}
}

func TestRemoveFromBucket(t *testing.T) {
	buck := NewBucket()
	id1 := NewRandomKademliaID()
	id2 := NewRandomKademliaID()
	id3 := NewRandomKademliaID()
	id4 := NewRandomKademliaID()

	conOne := NewContact(&id1, "localhost:8000")
	conTwo := NewContact(&id2, "localhost:8001")
	conThree := NewContact(&id3, "localhost:8002")
	conFour := NewContact(&id4, "localhost:8003")

	buck.AddContact(conOne)
	buck.AddContact(conTwo)
	buck.AddContact(conThree)
	buck.AddContact(conFour)
	buck.RemoveContact(conOne)
	if int64(buck.Len()) > 3 {
		t.Errorf("Cant remove contact!")
	}
}

func TestIsContactInBucket(t *testing.T) {
	buck := NewBucket()
	id1 := NewRandomKademliaID()
	id2 := NewRandomKademliaID()
	id3 := NewRandomKademliaID()
	id4 := NewRandomKademliaID()

	conOne := NewContact(&id1, "localhost:8000")
	conTwo := NewContact(&id2, "localhost:8001")
	conThree := NewContact(&id3, "localhost:8002")
	conFour := NewContact(&id4, "localhost:8003")

	buck.AddContact(conOne)
	buck.AddContact(conTwo)
	buck.AddContact(conThree)
	buck.AddContact(conFour)
	contactInBucket := buck.IsContactInBucket(conFour)
	if !contactInBucket {
		t.Errorf("Contact not in bucket!")
	}
}
