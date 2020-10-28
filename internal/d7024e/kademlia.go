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

	splitIP := strings.Split(myip, ".")
	splitIPBS := splitIP[:len(splitIP)-1]
	bootStrapSplitIP := append(splitIPBS, "3")
	bootStrapIP := strings.Join(bootStrapSplitIP, ".")

	bootStrapNodeStr := "2111111300000000000000000000123000000000"
	bootStrapKademID := NewKademliaID(&bootStrapNodeStr)
	bootStrapNode := NewContact(&bootStrapKademID, bootStrapIP)

	if splitIP[3] == "3" {
		return
	} else {

		i := 0
		for {

			pingRPC := NewRPC(Kademlia.network.routingTable.me, bootStrapIP, "PING", fmt.Sprint(i))
			Kademlia.network.SendMessage(pingRPC)
			time.Sleep(500 * time.Millisecond)
			ping, str := Kademlia.network.SendPINGMessage()
			if ping && str == fmt.Sprint(i) {
				LookUpContactRPC := NewRPC(bootStrapNode, Kademlia.network.routingTable.me.Address, "FINDNODE", "")
				Kademlia.network.SendMessage(LookUpContactRPC)
				break
			} else {
				fmt.Println("BOOOTSTRAPONNODENOTONLINE")
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
	closestJson := kademlia.LookupContact(&kademlia.network.routingTable.me)
	closest := kademlia.network.KTriples(closestJson)
	for i := 0; i < len(closest); i++ {
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)
		fmt.Println("Sent store to: " + closest[i].Address + " and data: " + data)
	}
}
