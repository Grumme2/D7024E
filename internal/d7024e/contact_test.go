package d7024e

// import (
// 	"testing"
// 	//"fmt"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCutCandidates(t *testing.T) {
// 	//mestr := "FFFFFFFF00000000000000000000000000000000"
// 	str1 := "FFFFFFFF00000000000000000000000000000000"
// 	str2 := "1111111100000000000000000000000000000000"

// 	//meid := NewKademliaID(&mestr)
// 	id1 := NewKademliaID(&str1)
// 	id2 := NewKademliaID(&str2)
// 	id3 := NewRandomKademliaID()
// 	id4 := NewRandomKademliaID()
// 	con1 := NewContact(&id1, "localhost")
// 	con2 := NewContact(&id2, "localhost")
// 	con3 := NewContact(&id3, "localhost")
// 	con4 := NewContact(&id4, "localhost")
// 	list := []Contact{con1,con2,con3,con4}
// 	canditates := ContactCandidates{list}
// 	canditates.CutContacts(2)
// 	assert.Equal(t, 2, canditates.Len())
// }