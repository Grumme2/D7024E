package d7024e

import (
	"testing"
	//"github.com/stretchr/testify/assert"
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
}