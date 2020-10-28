package d7024e

import (
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
	fmt.Println(myip)

	splitIP := strings.Split(myip, ".")
	x, splitIP := splitIP[len(splitIP)-1], splitIP[:len(splitIP)-1] //poping last element

	bootStrapSplitIP := append(splitIP, "3")
	bootStrapIP := strings.Join(bootStrapSplitIP, ".") // Bootstrap nodes iP address
	fmt.Println(bootStrapIP)
	bootStrapNodeStr := "2111111300000000000000000000123000000000" // hardcoded bootstrapnode ID
	bootStrapKademID := NewKademliaID(&bootStrapNodeStr)
	bootStrapNode := NewContact(&bootStrapKademID, bootStrapIP)
	fmt.Println(x)
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
			fmt.Println(message)
			if ping && message == fmt.Sprint(i) { // checks if ping gave a response

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
	//fmt.Println(shortlist)
	alreadyused := ContactCandidates{contacts: []Contact{}}
	str := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	id := NewKademliaID(&str)
	closestNode := NewContact(&id, "")
	closestNode.distance = &id
	less := shortlist.contacts[0].distance.Less(closestNode.distance)
	equal := shortlist.contacts[0].ID.Equals(target.ID)
	fmt.Println(less)
	fmt.Println(equal)
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(target.ID) {
		//fmt.Println("ENTERCHECK")
		closestNode = shortlist.contacts[0]
		fmt.Println(closestNode.ID)
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				findValueRPC := NewRPC(kademlia.network.routingTable.me, contact.Address, "FINDNODE", "")
				kademlia.network.SendMessage(findValueRPC)

				for j := 0; j < 11; j++ {
					fmt.Println("TEST")
					time.Sleep(500 * time.Millisecond)
					var data string
					data = kademlia.network.SendFindContactMessage()
					_ = data        //Ignores "data declared and not used" compilation error
					if data != "" { //If not undefined
						break //Exit for loop
					}
					fmt.Println(data)
					if j == 10 {
						fmt.Printf("hej")
						return "ERROR! Did not get response in time"
					}
				}
				data := kademlia.network.SendFindContactMessage()
				contacts := kademlia.network.KTriples(data)

				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				shortlist.Sort()
			}
		}
		shortlist.CutContacts(bucketSize)
	}
	KTrJson := kademlia.network.KTriplesJSON(shortlist.contacts)
	//fmt.Println(KTrJson)
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
	newkademid := NewKademliaID(&hash)
	closest := kademlia.network.routingTable.FindClosestContacts(&newkademid, alpha)
	//Add if statement to check if closest is empty
	closestNode := closest[0]
	shortlist := ContactCandidates{contacts: closest}
	alreadyused := ContactCandidates{contacts: []Contact{}}
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
				}

				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				shortlist.Sort()

			}
		}
		shortlist.CutContacts(bucketSize)
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
	closestjson := kademlia.LookupContact(&kademlia.network.routingTable.me)
	closest := kademlia.network.KTriples(closestjson)
	for i := 0; i < len(closest); i++ {
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)
		fmt.Println("Sent store to: " + closest[i].Address + " with data: " + data)
	}
}
