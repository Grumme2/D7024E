package d7024e

import (
	"fmt"
	"time"
)

type Kademlia struct {
	network Network
}

var alpha = 3

func (kademlia *Kademlia) LookupContact(target *Contact) string {
	shortlist := ContactCandidates{kademlia.network.routingTable.FindClosestContacts(target.ID, alpha)}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	closestNode := NewContact(NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"), "")
	closestNode.distance = NewKademliaID("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
	less := shortlist.contacts[0].distance.Less(closestNode.distance)
	equal := shortlist.contacts[0].ID.Equals(target.ID)
	fmt.Println(less)
	fmt.Println(equal)
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(target.ID) {
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
	return kademlia.network.KTriplesJSON(shortlist.contacts)

}

func in(a Contact, list []Contact) bool {
	for _, b := range list {
		if b.ID == a.ID {
			return true
		}
	}
	return false
}

func (kademlia *Kademlia) LookupData(hash string) string {
	kademlia.network.lookUpDataResponse = LookUpDataResponse{} //Resets LookUpDataResponse so we dont collect old results
	newkademid := NewKademliaID(hash)
	closest := kademlia.network.routingTable.FindClosestContacts(newkademid, alpha)
	closestNode := closest[0]
	shortlist := ContactCandidates{contacts: closest}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(newkademid) {
		closestNode = shortlist.contacts[0]
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				findValueRPC := NewRPC(kademlia.network.routingTable.me, contact.Address, "FINDVALUE", hash)
				kademlia.network.SendMessage(findValueRPC)

				for i := 0; i < 11; i++ {
					time.Sleep(500 * time.Millisecond)
					foundData, data := kademlia.network.SendFindDataMessage()
					_ = data                     //Ignores "data declared and not used" compilation error
					if foundData || !foundData { //If not undefined
						break //Exit for loop
					}
					if i == 10 {
						return "ERROR! Did not get response in time"
					}
				}

				foundData, data := kademlia.network.SendFindDataMessage()
				var contacts []Contact
				if foundData {
					return data
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
	return kademlia.network.KTriplesJSON(shortlist.contacts)
}

/*
func (kademlia *Kademlia) CreateNode(){
	network.Listen()


	myID := NewRandomKademliaID()
	me := NewContact(myID, myIP, myID)
	rt := NewRoutingTable(me)
}
*/

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

/*
func (kademlia *Kademlia) LookupContact(target *Contact) {
	rt.FindClosestContacts(target)

	//1. see if exists in routingtable
	//2. use rt.FindClosestContacts
	//3. see if exists in closests

}
*/

func (kademlia *Kademlia) Store(data string) {
	closestJson := kademlia.LookupContact(&kademlia.network.routingTable.me)
	closest := kademlia.network.KTriples(closestJson)
	for i := 0; i < len(closest); i++ {
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)
	}
}
