package d7024e

type Kademlia struct {
	Rt      RoutingTable
	network Network
}

var alpha = 3

func (kademlia *Kademlia) LookupContact(target *Contact) []Contact {
	closest := kademlia.Rt.FindClosestContacts(target.ID, alpha)
	closestNode := closest[0]
	shortlist := ContactCandidates{contacts: closest}
	alreadyused := ContactCandidates{contacts: []Contact{}}
	for shortlist.contacts[0].distance.Less(closestNode.distance) && !shortlist.contacts[0].ID.Equals(target.ID) {
		closestNode := shortlist.contacts[0]
		for i := 0; i < 3; i++ {
			contact := shortlist.contacts[i]
			if !in(contact, alreadyused.contacts) {
				contacts := contact.Kademlia.network.SendFindContactMessage(target)
				alreadyused.Append([]Contact{contact})
				shortlist.Append(contacts)
				shortlist.Sort()
				bucket := kademlia.Rt.buckets[kademlia.Rt.getBucketIndex(shortlist.contacts[i].ID)]
				addToBucket(contact, *bucket)
			}

		}
		shortlist.CutContacts(bucketSize)
	}
	return shortlist.contacts

}

func in(a Contact, list []Contact) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func addToBucket(contact Contact, buck bucket) {
	if (buck.Len() < bucketSize) && buck.In(contact) {
		buck.AddContact(contact)
	} else {
		//sping(buck.list.Front())
		buck.AddContact(contact)
	}

}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
