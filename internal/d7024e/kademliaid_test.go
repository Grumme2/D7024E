package d7024e

import(
	"encoding/hex"
	"testing"
	//"fmt"
	"github.com/stretchr/testify/assert"

)


func TestNewKademliaID (t *testing.T) {
	kademID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	// length := len(kademID)
	// fmt.Println(kademID)

	// if length==IDLength{
		// t.Errorf("NewKademliaID not initiated correctly!")
	// }
	str :=  hex.EncodeToString(kademID[0:IDLength])
	assert.NotNil(t, kademID)
	assert.Equal(t, str , "ffffffff00000000000000000000000000000000")
	assert.Equal(t, len(str), 2*IDLength)
}
func TestNewRandomKademliaID (t *testing.T) {
	kademID := NewRandomKademliaID()
	if len(hex.EncodeToString(kademID[0:IDLength]))==IDLength {
		t.Errorf("not initiated correctly!")
	}
}

func TestLess(t *testing.T){
	kademID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	kademID2 := NewKademliaID("FFFFFFF000000000000000000000000000000000")

	assert.Equal(t,kademID.Less(kademID2), false)
	assert.Equal(t,kademID2.Less(kademID), true)
}


func TestEqual(t *testing.T){
	kademID := NewKademliaID("FFFFFFFF00000000000000000000000000000000")
	kademID2 := NewKademliaID("FFFFFFF000000000000000000000000000000000")

	assert.Equal(t,kademID.Equals(kademID), true)
	assert.Equal(t,kademID.Equals(kademID2), false)
}