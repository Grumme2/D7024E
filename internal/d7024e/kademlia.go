package d7024e

import (
	"encoding/hex"
	"fmt"
	"strings"
	"time"
)

type Kademlia struct {
	network *Network
}

func NewKademlia(network *Network) Kademlia {
	return Kademlia{network}
}

var alpha = 3

func (Kademlia *Kademlia) JoinNetwork() {
	myip := GetLocalIP()

	splitIP := strings.Split(myip, ".")
	x, splitIP := splitIP[len(splitIP)-1], splitIP[:len(splitIP)-1] //poping last element

	bootStrapSplitIP := append(splitIP, "3")
	bootStrapIP := strings.Join(bootStrapSplitIP, ".")             // Bootstrap nodes iP address
	bootStrapNodeStr := "2111111300000000000000000000123000000000" // hardcoded bootstrapnode ID
	bootStrapKademID := NewKademliaID(bootStrapNodeStr)
	bootStrapNode := NewContact(&bootStrapKademID, bootStrapIP)

	if x == "3" { // if Bootstrap node nothing needs to be done
		fmt.Println("bootstrapnode")
		return
	} else {
		i := 0
		for {

			pingRPC := NewRPC(Kademlia.network.routingTable.me, bootStrapIP, "PING", fmt.Sprint(i))
			Kademlia.network.SendMessage(pingRPC)
			time.Sleep(500 * time.Millisecond)
			ping, message := Kademlia.network.SendPINGMessage() // gets ping result
			if ping && message == fmt.Sprint(i) {               // checks if ping gave a response

				// LookUpContactRPC := NewRPC(bootStrapNode, Kademlia.network.routingTable.me.Address, "FINDNODE", "") // lookup contact rpc should just call lookupcontact
				// Kademlia.network.SendMessage(LookUpContactRPC)
				//
				Kademlia.network.routingTable.AddContact(bootStrapNode)
				Kademlia.LookupContact(&Kademlia.network.routingTable.me) // Broken right now so are just doing the rpc call

				break // breaks when node has joined network

			} else {
				fmt.Println("BOOOTSTRAP_NODE_NOT_ONLINE")
			}

			i++
		}
	}
}

func (kademlia *Kademlia) LookupContact(target *Contact) string {
	shortlist := ContactCandidates{kademlia.network.routingTable.FindClosestContacts(target.ID, alpha)}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	id := KademliaID{}
	decoded, _ := hex.DecodeString("FF")

	for i := 0; i < IDLength; i++ {
		id[i] = decoded[0]
	}
	closestNode := NewContact(&id, "")
	closestNode.distance = &id

	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(target.ID) {

		closestNode = shortlist.contacts[0]

		for i := 0; i < 3; i++ {
			if len(shortlist.contacts) <= i {
				break
			}
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				findNodeRPC := NewRPC(kademlia.network.routingTable.me, contact.Address, "FINDNODE", "")
				kademlia.network.SendMessage(findNodeRPC)

				for j := 0; j < 11; j++ {

					time.Sleep(500 * time.Millisecond)
					var data string
					data = kademlia.network.SendFindContactMessage()
					_ = data        //Ignores "data declared and not used" compilation error
					if data != "" { //If not undefined
						break //Exit for loop
					}

					if j == 10 {
						fmt.Printf("hej")
						return "ERROR! Did not get response in time"
					}
				}
				data := kademlia.network.SendFindContactMessage()
				contacts := kademlia.network.KTriples(data)
				for k := 0; k < len(contacts); k++ {
					contacts[k].CalcDistance(kademlia.network.routingTable.me.ID)

				}
				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				if len(shortlist.contacts) >= 2 {
					for i := 0; i < len(shortlist.contacts); i++ {
						for j := i + 1; j < len(shortlist.contacts); j++ {
							if shortlist.contacts[j].Less(&shortlist.contacts[i]) {
								temp := shortlist.contacts[i]
								shortlist.contacts[i] = shortlist.contacts[j]
								shortlist.contacts[j] = temp
							}
						}
					}
				}

			}
		}
		if len(shortlist.contacts) > bucketSize {
			shortlist.CutContacts(bucketSize)
		}

	}
	KTrJson := kademlia.network.KTriplesJSON(shortlist.contacts)
	fmt.Println(KTrJson)
	fmt.Println(kademlia.network.routingTable.me)
	return KTrJson

}

func in(a Contact, list []Contact) bool {
	for _, b := range list {
		if b.ID == a.ID {
			return true
		}
	}
	return false
}

func (kademlia *Kademlia) LookupData(hash string) (bool, string, Contact) {
	kademlia.network.lookUpDataResponse = LookUpDataResponse{} //Resets LookUpDataResponse so we dont collect old results
	newkademid := NewKademliaID(hash)

	shortlist := ContactCandidates{kademlia.network.routingTable.FindClosestContacts(&newkademid, alpha)}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	id := KademliaID{}
	decoded, _ := hex.DecodeString("FF")

	for i := 0; i < IDLength; i++ {
		id[i] = decoded[0]
	}
	closestNode := NewContact(&id, "")
	closestNode.distance = &id
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(&newkademid) {
		closestNode = shortlist.contacts[0]
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				findValueRPC := NewRPC(kademlia.network.routingTable.me, contact.Address, "FINDVALUE", hash)
				kademlia.network.SendMessage(findValueRPC)

				for j := 0; j < 11; j++ {
					time.Sleep(500 * time.Millisecond)
					foundData, data, node := kademlia.network.SendFindDataMessage()
					_ = data                     //Ignores "data declared and not used" compilation error
					_ = node                     //Ignores "data declared and not used" compilation error
					if foundData || !foundData { //If not undefined
						break //Exit for loop
					}
				}

				foundData, data, node := kademlia.network.SendFindDataMessage()
				var contacts []Contact
				if foundData {
					return true, data, node
				} else {
					contacts = kademlia.network.KTriples(data)
					for k := 0; k < len(contacts); k++ {
						contacts[k].CalcDistance(kademlia.network.routingTable.me.ID)

					}
				}

				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				if len(shortlist.contacts) >= 2 {
					for i := 0; i < len(shortlist.contacts); i++ {
						for j := i + 1; j < len(shortlist.contacts); j++ {
							if shortlist.contacts[j].Less(&shortlist.contacts[i]) {
								temp := shortlist.contacts[i]
								shortlist.contacts[i] = shortlist.contacts[j]
								shortlist.contacts[j] = temp
							}
						}
					}
				}

			}
		}
		if len(shortlist.contacts) > bucketSize {
			shortlist.CutContacts(bucketSize)
		}
	}
	return false, kademlia.network.KTriplesJSON(shortlist.contacts), kademlia.network.routingTable.me
}

/*
func (kademlia *Kademlia) JoinNetwork(target *Contact) {
	//Generate new kademlia for self if none exists
	if (rt.me == nil) {
		myIP := 127.0.0.1 //how get own ip???????????????????? am dumb
		myID := NewRandomKademliaID()
		rt.me := NewContact(myID, myIP, myID)
	}

	rt.AddContact(target)                     //Adds the target to the correct k-bucket
	LookupContact(rt.me)                      //Node lookup on self
	closetsContacts := rt.FindClosestContacts //Get closests contacts

	//Refresh all contacts except the closets neighbour (which is index 0 in the array)
	for i := range closetsContacts {
		if i != 0 {
			LookupContact(closetsContacts[i])
		}
	}
}
*/

func (kademlia *Kademlia) Store(data string) {
	closestJson := kademlia.LookupContact(&kademlia.network.routingTable.me)
	closest := kademlia.network.KTriples(closestJson)
	for i := 0; i < len(closest); i++ {
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)
		fmt.Println("Sent store to: " + closest[i].Address + " with data: " + data)
	}
}
