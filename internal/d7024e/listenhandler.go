package d7024e

import(
	"fmt"
	"time"
	//"net"
)

func (network *Network) ListenHandler() {
	for {
		receivedData := <- network.internal
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
}