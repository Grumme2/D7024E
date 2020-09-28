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

func (network *Network) SendFindContactMessage(contact *Contact) Contact {
	if contact.ID == network.kadem.Rt.me.ID {
		return *contact
	}
	contacts := network.kadem.Rt.FindClosestContacts(contact.ID, 1)
	return contacts[0].Kademlia.network.SendFindContactMessage(contact)
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
