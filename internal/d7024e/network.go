package d7024e

import (
	"container/list"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type Network struct {
	routingTable          *RoutingTable
	awaitingResponseList  *list.List
	lookUpDataResponse    LookUpDataResponse
	lookUpContactResponse LookUpContactResponse
	pingResponse          PINGResponse
}

type LookUpDataResponse struct {
	DataFound bool
	Data      string
	Node      Contact
}

type LookUpContactResponse struct {
	data string
}

type PINGResponse struct {
	pong bool
	data string
}

type AwaitingResponseObject struct {
	timestamp int64
	oldNode   Contact
	newNode   Contact
}

func NewNetwork(rt *RoutingTable) Network {
	emptyLookUpData := LookUpDataResponse{}
	emptyLookUpContact := LookUpContactResponse{}
	emptyPINGresponse := PINGResponse{}
	return Network{rt, list.New(), emptyLookUpData, emptyLookUpContact, emptyPINGresponse}
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

	time.Sleep(2 * time.Second) //Checks every two seconds
	network.CheckNodesAwaitingResponse()
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
		network.ListenHandler(receivedData, connection, addr)
		_ = err
	}
}

func (network *Network) SendMessage(message RPC) {
	CONNECT := message.TargetAddress + ":8000" //Hardcoded port

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	for {
		data := []byte(JSONEncode(message))
		fmt.Println("SENT:\n", string(data))
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return
		}
		return
	}
}

func (network *Network) ListenHandler(receivedData []byte, connection *net.UDPConn, addr *net.UDPAddr) {
	decodedData := JSONDecode(receivedData)

	fmt.Println("RECEIVED:\n", string(receivedData))

	//True if contact already is in bucket
	if network.routingTable.buckets[network.routingTable.GetBucketIndex(decodedData.Sender.ID)].IsContactInBucket(decodedData.Sender) {
		network.routingTable.AddContact(decodedData.Sender) //Move contact to start of bucket
	} else if network.routingTable.IsBucketFull(decodedData.Sender.ID) {
		//If bucket is full, the node pings the contact at the tail of the buckets list
		//If previously mentioned contact fails to respond in x amount of time, it is dropped from the list and the new contact is added at the head
		bucketIndex := network.routingTable.GetBucketIndex(decodedData.Sender.ID)
		tailContact := network.routingTable.buckets[bucketIndex].list.Back().Value.(Contact) //Vet ej om detta faktiskt stämmer
		currentTime := time.Now().Unix()

		awaitingResponseData := AwaitingResponseObject{currentTime, tailContact, decodedData.Sender}
		network.awaitingResponseList.PushFront(awaitingResponseData)
		pingRPC := NewRPC(network.routingTable.me, tailContact.Address, "PING", "")
		network.SendMessage(pingRPC)
	} else {
		network.routingTable.AddContact(decodedData.Sender) //Adds contact to start of the bucket
	}

	responseType := "UNDEFINED"
	responseContent := "defaultNetworkResponse"

	switch decodedData.MessageType {
	case "PING":
		responseType = "PONG"
		responseContent = decodedData.Content
	case "OK":
		responseType = "NONE"
	case "STORE":
		key := network.AddToStore(decodedData.Content)
		responseType = "OK"
		responseContent = key
	case "FINDVALUE":
		dataFound, data := network.LookForData(decodedData.Content)
		if dataFound {
			responseType = "FINDVALUE_RESPONSE"
			lookupResponse := LookUpDataResponse{true, data, network.routingTable.me}
			responseContent = network.JSONEncodeLookUpDataResponse(lookupResponse)
		} else {
			responseType = "FINDVALUE_RESPONSE"
			closest := network.routingTable.FindClosestContacts(network.routingTable.me.ID, bucketSize)
			closestEncoded := network.KTriplesJSON(closest)
			lookupResponse := LookUpDataResponse{false, closestEncoded, network.routingTable.me}
			responseContent = network.JSONEncodeLookUpDataResponse(lookupResponse)
		}
	case "FINDVALUE_RESPONSE":
		var data = network.JSONDecodeLookUpDataResponse(decodedData.Content)
		network.lookUpDataResponse = data
		responseType = "NONE"
	case "FINDNODE":
		responseType = "FINDNODE_RESPONSE"
		closest := network.routingTable.FindClosestContacts(network.routingTable.me.ID, bucketSize)
		//fmt.Println(closest)
		responseContent = network.KTriplesJSON(closest)
		//fmt.Println(closestEncoded)

	case "FINDNODE_RESPONSE":
		fmt.Println("FFFFFFFFFFUUUUUUUUCCCCCCCCKKKKKKKKKK")
		network.lookUpContactResponse = LookUpContactResponse{decodedData.Content}
		fmt.Println(network.lookUpContactResponse)
		responseType = "NONE"

	case "PONG":
		network.pingResponse = PINGResponse{true, decodedData.Content}
		responseType = "NONE"
	}

	if responseType != "NONE" && responseType != "UNDEFINED" {
		responseRPC := NewRPC(network.routingTable.me, decodedData.Sender.Address, responseType, responseContent)
		network.SendMessage(responseRPC)

	} else { //No response if the messagetype is NONE or UNDEFINED
		fmt.Println("Received 'OK', 'PONG' or 'UNDEFINED' message. Conversation done.")
	}
}

func (network *Network) AddToStore(message string) string {
	hxMsg := network.MakeHash(message)
	network.routingTable.me.KeyValueStore[hxMsg] = message
	return hxMsg
}

func (network *Network) LookForData(hash string) (bool, string) {
	for key, element := range network.routingTable.me.KeyValueStore {
		if key == hash {
			fmt.Println("LookForData found element: " + element)
			return true, element
		}
	}
	return false, ""
}

func (network *Network) MakeHash(message string) string {
	hx := hex.EncodeToString([]byte(message))
	return hx
}

func (network *Network) storeRPC(message RPC) {
	hash := network.MakeHash(message.Content)
	fmt.Printf(hash)
	network.SendMessage(message)
}

func (network *Network) JSONEncodeLookUpDataResponse(unencodedResponse LookUpDataResponse) string {
	encoded, err := json.Marshal(unencodedResponse)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	encodedString := string(encoded)
	return encodedString
}

func (network *Network) JSONDecodeLookUpDataResponse(encodedString string) LookUpDataResponse {
	var unencoded LookUpDataResponse
	err := json.Unmarshal([]byte(encodedString), &unencoded)
	if err != nil {
		fmt.Println(err)
	}
	return unencoded
}

func (network *Network) KTriplesJSON(KClosest []Contact) string {
	contactsJSON, err := json.Marshal(KClosest)
	if err != nil {
		fmt.Println(err)
		return "ERROR"
	}
	contactsStr := string(contactsJSON)
	return contactsStr
}

func (network *Network) KTriples(KClosest string) []Contact {
	var contacts []Contact
	err := json.Unmarshal([]byte(KClosest), &contacts)
	if err != nil {
		fmt.Println(err)
		//return "ERROR"
	}
	fmt.Println(contacts)
	return contacts
}

func (network *Network) SendFindContactMessage() string {
	// contacts := network.routingTable.FindClosestContacts(contact.ID, bucketSize)
	return network.lookUpContactResponse.data
}

func (network *Network) SendFindDataMessage() (bool, string, Contact) {
	return network.lookUpDataResponse.DataFound, network.lookUpDataResponse.Data, network.lookUpDataResponse.Node
}

func (network *Network) SendPINGMessage() (bool, string) {
	return network.pingResponse.pong, network.pingResponse.data
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
