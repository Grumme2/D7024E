package d7024e

type Kademlia struct {
	Rt      RoutingTable
	network Network
}

var alpha = 3

func (kademlia *Kademlia) LookupContact(target *Contact) *Contact {
	contacts := kademlia.Rt.FindClosestContacts(target.ID, alpha)
	for i := range contacts {
		contact := contacts[i].Kademlia.network.SendFindContactMessage(target)
		if contact == *target {
			return &contact
		}
	}
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
