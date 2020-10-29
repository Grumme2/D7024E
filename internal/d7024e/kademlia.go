package d7024e

import (
	"fmt"
	"time"
	"container/list"
)

type Kademlia struct {
	network *Network
}

func NewKademlia(network *Network) Kademlia {
	return Kademlia{network}
}

var alpha = 3

func (kademlia *Kademlia) LookupContact(target *Contact) string {
	kademlia.network.lookUpContactResponse.Data = "undefined" //Resets LookUpContactResponse so we dont collect old results

	//Hämtar tre närmsta i ordning
	contactList := kademlia.network.routingTable.FindClosestContacts(target.ID, alpha)
	//[]Contact contactedNodes := network.routingTable.me
	contactedNodes := list.New()
	closestNode := contactList[0] //First contact given by FindClosestContacts is always the closest
	newClosestNodeFound := true

	for (newClosestNodeFound) {
		newClosestNodeFound = false
		localContactList := contactList
		
		for i := len(localContactList); i > 0; i-- {
			nodeAlreadyContacted := false

			//Checks if the current contact exists in the contactedNodes list
			for e := contactedNodes.Front(); e != nil; e = e.Next() {
				if localContactList[i].ID == e.Value.(Contact).ID {
					contactList = append(contactList[:i], contactList[i+1:]...) //Removes the i:th element from contactList
					nodeAlreadyContacted = true
				}
			}

			if (!nodeAlreadyContacted) {
				//TODO: only run if not already contacted !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
				findValueRPC := NewRPC(kademlia.network.routingTable.me, localContactList[i].Address, "FINDNODE", "LookupContact")
				kademlia.network.SendMessage(findValueRPC)
				time.Sleep(100 * time.Millisecond) //Give it some time to respond
				response := kademlia.network.GetFindContactResponse()
				if (response != "undefined"){
					responseContacts := kademlia.network.KTriples(response)
					for j := 0; j < len(responseContacts); j++ {
						contactList[len(contactList)+1] = responseContacts[j]
						if (responseContacts[j].Less(&closestNode)){
							closestNode = responseContacts[j]
							newClosestNodeFound = true
						}
					}
				} else {
					contactList = append(contactList[:i], contactList[i+1:]...) //Removes the i:th element from contactList
				}
				contactedNodes.PushBack(localContactList[i])
			}
		}
	}

	//Manually sorts contacts
	for i := 0; i < len(contactList); i++ {
		for j := i+1; j < len(contactList); j++ {
			if (contactList[j].Less(&contactList[i])){
				temp := contactList[i]
				contactList[i] = contactList[j]
				contactList[j] = temp
			}
		}
	}

	return kademlia.network.KTriplesJSON(contactList[:3])
}

/*
func (kademlia *Kademlia) LookupContact(target *Contact) string {
	kademlia.network.lookUpContactResponse = LookUpContactResponse{} //Resets LookUpContactResponse so we dont collect old results
	//Hämtar 3 närmsta contacts
	shortlist := ContactCandidates{kademlia.network.routingTable.FindClosestContacts(target.ID, alpha)}
	//fmt.Println(shortlist)
	//Lista med redan använda contacts, varför inte bara ta bort ur listan eller loopa?
	alreadyused := ContactCandidates{contacts: []Contact{}}
	str := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF" //Maxdistans
	id := NewKademliaID(&str)
	closestNode := NewContact(&id, "")
	closestNode.distance = &id
	less := shortlist.contacts[0].distance.Less(closestNode.distance) //Bool: Är första kontakten mindre än maxvärdet
	equal := shortlist.contacts[0].ID.Equals(target.ID) //Bool: Är första kontakten lika med maxvärdet
	fmt.Println(less) //Är true
	fmt.Println(equal) //Är true!!!!!!!!!! Vilket gör att vi inte går in i loopen
	for true { //While (mindre än max && ej lika med max)
		fmt.Println("ENTERCHECK")
		closestNode = shortlist.contacts[0] //Närmsta noden
		fmt.Println(closestNode.ID)
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			fmt.Println("Lookuploop IP: " + contact.Address)
			if !in(contact, alreadyused.contacts) { //Om contact inte finns i redan använda contacts
				//Skicka RPC
				findValueRPC := NewRPC(kademlia.network.routingTable.me, contact.Address, "FINDNODE", "")
				kademlia.network.SendMessage(findValueRPC)

				for j := 0; j < 6; j++ {
					fmt.Println("TEST")
					time.Sleep(500 * time.Millisecond)
					var data string
					data = kademlia.network.SendFindContactMessage() //Hämtar data från struct i network
					_ = data        //Ignores "data declared and not used" compilation error
					if data != "" { //If not undefined
						break //Exit for loop
					}
					fmt.Println(data)
					if j == 5 {
						fmt.Printf("ERROR!!!!!!!!!")
						return "ERROR! Did not get response in time"
					}
				}
				data := kademlia.network.SendFindContactMessage()
				contacts := kademlia.network.KTriples(data)

				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				//shortlist.Sort() //KRASH
			}
		}
		shortlist.CutContacts(bucketSize) //Cut till 20? Fungerar detta???? //KRASHAR
	}
	KTrJson := kademlia.network.KTriplesJSON(shortlist.contacts)
	//fmt.Println(KTrJson)
	return KTrJson

}
*/

func in(a Contact, list []Contact) bool {
	for _, b := range list {
		if b.ID == a.ID {
			return true
		}
	}
	return false
}

func (kademlia *Kademlia) LookupData(hash string) (bool, string, Contact) {
	return true, "hej", kademlia.network.routingTable.me
}

/*
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
				fmt.Println("LookupData sent to address: " + contact.Address)
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

func (kademlia *Kademlia) Store(data string) {
	//Closests JSON innehåller alltid endast en contact !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
	closestjson := kademlia.LookupContact(&kademlia.network.routingTable.me)
	fmt.Println("closest json")
	fmt.Println(closestjson)
	closest := kademlia.network.KTriples(closestjson)
	fmt.Println("CLOSEST BELOW:")
	fmt.Println(closest)
	fmt.Println("Store: Length of closests: " + string(len(closest)))
	fmt.Println(len(closest))
	for i := 0; i < len(closest); i++ {
		fmt.Println("Sending store to IP: " + closest[i].Address)
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)
		fmt.Println("Sent store to: " + closest[i].Address + " with data: " + data)
	}
}
