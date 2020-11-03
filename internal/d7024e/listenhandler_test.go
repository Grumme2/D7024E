package d7024e
// import (
// 	"fmt"
// 	"testing"
// 	"time"

// 	//"reflect"
// 	//"github.com/stretchr/testify/assert"
// )

// func TestStoreRPC(t *testing.T)  {
// 	//id1 := NewRandomKademliaID()
// 	sender := NewRandomKademliaID()
// 	time.Sleep(2 * time.Millisecond)
// 	recivier := NewRandomKademliaID()
// 	time.Sleep(2 * time.Millisecond)
// 	rtContactId := NewRandomKademliaID()

// 	me := NewContact(&sender, "localhost")
// 	//rtContact := NewContact(&rtContactId, "localhost")
// 	fmt.Println(rtContactId)
// 	fmt.Println(recivier)
// 	fmt.Println(sender)
// 	rt := NewRoutingTable(me)
// 	network := NewNetwork(rt)
// 	network.testing = true
// 	kademlia := NewKademlia(&network)
// 	kademlia.network.routingTable.AddContact(NewContact(&recivier, "localhost"))
// 	storeRPC := NewRPC(me, "localhost","STORE", "hej")
// 	storeRPCJson := JSONEncode(storeRPC)
// 	go kademlia.network.ListenHandler()
// 	msg := <- kademlia.network.external
	
// 	kademlia.network.internal <- storeRPCJson
	
// 	msg2 := <- kademlia.network.external
// 	fmt.Print("msg: " + string(msg))
// 	fmt.Print("msg2: " + string(msg2) )
	



// }