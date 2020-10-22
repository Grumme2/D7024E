package d7024e

import (
	//"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
//	"strconv"
)

func TestJSONDecode (t *testing.T) {
	senderCon := NewContact(NewRandomKademliaID(), "localhost:8000")
	RPCPreUnmarshal := NewRPC(senderCon, "localhost:8001", "PING", "")
	encodedRPC := JSONEncode(RPCPreUnmarshal)
	decodedRPC := JSONDecode(encodedRPC)
	assert.Equal(t, RPCPreUnmarshal.MessageType,decodedRPC.MessageType)
}


