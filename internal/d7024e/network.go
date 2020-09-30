package d7024e

type Network struct {
	adress string
	Id     KademliaID
	kadem  *Kademlia
}

func Listen(ip string, port int) {
	// TODO
}

func (network *Network) SendPingMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindContactMessage(contact *Contact) []Contact {
	contacts := network.kadem.Rt.FindClosestContacts(contact.ID, k)
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
