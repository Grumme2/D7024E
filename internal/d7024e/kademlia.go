package d7024e

type Kademlia struct {
	rt      RoutingTable
	network Network
}

func (kademlia *Kademlia) JoinNetwork(target *Contact) {

	if rt.me == nil {
		//generate kademliaId etc for me
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

func (kademlia *Kademlia) LookupContact(target *Contact) {
	rt.FindClosestContacts(target)

	//1. see if exists in routingtable
	//2. use rt.FindClosestContacts
	//3. see if exists in closests
}

func (kademlia *Kademlia) LookupData(hash string) {
	// TODO
}

func (kademlia *Kademlia) Store(data []byte) {
	// TODO
}
