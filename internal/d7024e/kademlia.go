package d7024e

type Kademlia struct {
	network Network
}

var alpha = 3

func (kademlia *Kademlia) LookupContact(target *Contact) []Contact {
	closest := kademlia.network.routingTable.FindClosestContacts(target.ID, alpha)
	closestNode := closest[0]
	shortlist := ContactCandidates{contacts: closest}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(target.ID) {
		closestNode = shortlist.contacts[0]
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				contacts := kademlia.network.SendFindContactMessage(target)
				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				shortlist.Sort()
			}

		}
		shortlist.CutContacts(bucketSize)
	}
	return shortlist.contacts

}

func in(a Contact, list []Contact) bool {
	for _, b := range list {
		if b.ID == a.ID {
			return true
		}
	}
	return false
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
	closest := kademlia.LookupContact(&kademlia.network.routingTable.me)
	for i := 0; i < len(closest); i++ {
		rpc := NewRPC(kademlia.network.routingTable.me, closest[i].Address, "STORE", data)
		kademlia.network.storeRPC(rpc)	
	}
}
