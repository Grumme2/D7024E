package d7024e

import (
	"testing"
)

func TestAddToBucket(t *testing.T) {
	buck := NewBucket()
	conOne := NewContact(NewRandomKademliaID(), "localhost:8000")
	conTwo := NewContact(NewRandomKademliaID(), "localhost:8001")
	conThree := NewContact(NewRandomKademliaID(), "localhost:8002")
	conFour := NewContact(NewRandomKademliaID(), "localhost:8003")
	conFive := NewContact(NewRandomKademliaID(), "localhost:8004")
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
	conOne := NewContact(NewRandomKademliaID(), "localhost:8000")
	conTwo := NewContact(NewRandomKademliaID(), "localhost:8001")
	conThree := NewContact(NewRandomKademliaID(), "localhost:8002")
	conFour := NewContact(NewRandomKademliaID(), "localhost:8003")
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
	conOne := NewContact(NewRandomKademliaID(), "localhost:8000")
	conTwo := NewContact(NewRandomKademliaID(), "localhost:8001")
	conThree := NewContact(NewRandomKademliaID(), "localhost:8002")
	conFour := NewContact(NewRandomKademliaID(), "localhost:8003")
	buck.AddContact(conOne)
	buck.AddContact(conTwo)
	buck.AddContact(conThree)
	buck.AddContact(conFour)
	contactInBucket := buck.IsContactInBucket(conFour)
	if !contactInBucket {
		t.Errorf("Contact not in bucket!")
	}
}
