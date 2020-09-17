package d7024e

type Kademlia struct {
	rt RoutingTable
	network Network
}

func (node *Node) JoinNetwork(target *Contact) {
	//Must already have contact to a node N in the network
	rt.AddContact(target) //1. Insert N into k-bucket
	LookupContact(rt.me) //2. Node lookup on its own node ID
	//3a. refresh all k-buckets further away than its closets neighbour
	//3b. during the refreshes this node populates its own k-buckets and 
	//bucket.GetContactAndCalcDistance()
	//rt.FindClosestContacts()

	//4. inserts itself into other nodes k-buckets as necessary
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
