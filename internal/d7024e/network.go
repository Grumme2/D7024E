package d7024e

import (
	"container/list"
	"encoding/hex"
	"fmt"
	"net"
	"time"
)

type Network struct {
	routingTable         *RoutingTable
	awaitingResponseList *list.List
}

type AwaitingResponseObject struct {
	timestamp int64
	oldNode   Contact
	newNode   Contact
}

func NewNetwork(rt *RoutingTable) Network {
	return Network{rt, list.New()}
}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

//Checks the awaitingResponseOldNode and awaitingResponseNewNode and checks if the old node has responded.
//If it hasn't the old node is removed and the new one is added. Otherwise nothing happens.
func (network *Network) CheckNodesAwaitingResponse() {
	currentTime := time.Now().Unix()

	for e := network.awaitingResponseList.Front(); e != nil; e = e.Next() {
		nodeTimestamp := e.Value.(AwaitingResponseObject).timestamp
		fmt.Println(nodeTimestamp)
		fmt.Println(currentTime)
		fmt.Println(currentTime - nodeTimestamp)
		fmt.Println(e.Value.(AwaitingResponseObject).oldNode)
		fmt.Println(e.Value.(AwaitingResponseObject).newNode)
		if (currentTime - nodeTimestamp) >= 5 { //If 5 seconds or more have passed
			network.routingTable.RemoveContact(e.Value.(AwaitingResponseObject).oldNode)
			network.routingTable.AddContact(e.Value.(AwaitingResponseObject).newNode)
			fmt.Println("remove bucket")
		}
		//Else we do nothing, the old node remains and the new node is not added
	}
}

func (network *Network) TestCreateAwaitingReponseObjects() {
	currentTime := time.Now().Unix()
	awaitingResponseData := AwaitingResponseObject{currentTime, network.routingTable.me, NewContact(network.routingTable.me.ID, "hehehe")}
	network.awaitingResponseList.PushFront(awaitingResponseData)
}

func (network *Network) Listen() {
	PORT := ":8000" //Hardcoded port

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		receivedData := buffer[0:n]
		decodedData := JSONDecode(receivedData)

		fmt.Printf("RECEIVED: %s\n", string(receivedData))

		//True if contact already is in bucket
		if network.routingTable.buckets[network.routingTable.getBucketIndex(decodedData.Sender.ID)].IsContactInBucket(decodedData.Sender) {
			network.routingTable.AddContact(decodedData.Sender) //Move contact to start of bucket
		} else if network.routingTable.isBucketFull(decodedData.Sender.ID) {
			//If bucket is full, the node pings the contact at the tail of the buckets list
			//If previously mentioned contact fails to respond in x amount of time, it is dropped from the list and the new contact is added at the head
			bucketIndex := network.routingTable.getBucketIndex(decodedData.Sender.ID)
			tailContact := network.routingTable.buckets[bucketIndex].list.Back().Value.(Contact) //Vet ej om detta faktiskt stämmer
			currentTime := time.Now().Unix()

			awaitingResponseData := AwaitingResponseObject{currentTime, tailContact, decodedData.Sender}
			network.awaitingResponseList.PushFront(awaitingResponseData)
			pingRPC := NewRPC(network.routingTable.me, tailContact.Address, "PING", "")
			network.SendPingMessage(pingRPC)

		} else {
			network.routingTable.AddContact(decodedData.Sender) //Adds contact to start of the bucket
		}

		if decodedData.MessageType != "NONE" && decodedData.MessageType != "UNDEFINED" {
			responseType := "UNDEFINED"
			responseContent := "defaultNetworkResponse"

			switch decodedData.MessageType {
			case "PING":
				responseType = "PONG"
			case "PONG":
				responseType = "OK"
			case "OK":
				responseType = "NONE"
			}

			responseRPC := NewRPC(network.routingTable.me, decodedData.Sender.Address, responseType, responseContent)
			responseData := JSONEncode(responseRPC)
			response := []byte(responseData)

			fmt.Printf("SENT: %s\n", string(response))

			_, err = connection.WriteToUDP(response, addr)
			if err != nil {
				fmt.Println(err)
				return
			}
		} else { //No response if the messagetype is NONE or UNDEFINED
			fmt.Println("Received 'OK' or 'UNDEFINED' message. Will not respond.")
		}
	}
}

func (network *Network) SendMessage(message RPC) bool {
	CONNECT := message.TargetAddress + ":8000" //Hardcoded port

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer c.Close()

	for {
		data := []byte(JSONEncode(message))
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return false
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return false
		}

		fmt.Printf("RECEIVED: %s\n", string(buffer[0:n]))
		return true
	}
}

func (network *Network) AddToStore(message string) {
	hxMsg := hex.EncodeToString([]byte(message))
	KeyValueStore[hxMsg] = message
}

func (network *Network) LookForData(hash string) string {
	for key, element := range KeyValueStore {
		if key == hash {
			return element
		} else {
			continue
		}
	}
	return ""
}

func (network *Network) SendFindContactMessage(contact *Contact) []Contact {
	contacts := network.routingTable.FindClosestContacts(contact.ID, bucketSize)
	return contacts
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
