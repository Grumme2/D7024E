package d7024e

import (
	"fmt"
	"net"
)

type Network struct {
	routingTable *RoutingTable
}

func NewNetwork() Network {
	return Network{}
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

func (network *Network) Listen(rt RoutingTable) {
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

		//For every recieved communication, update the bucket corresponding to the sender.	
		//If contact already exists, move it to the start of the bucket.

		//True if contact already is in bucket
		if(rt.buckets[rt.getBucketIndex(decodedData.Sender.ID)].IsContactInBucket(decodedData.Sender)){
			rt.AddContact(decodedData.Sender) //Move contact to start of bucket
		} else if (rt.isBucketFull(decodedData.Sender.ID)){
			//If bucket is full, the node pings the contact at the tail of the buckets list
			//If previously mentioned contact fails to respond in x amount of time, it is dropped from the list and the new contact is added at the head
			//Otherwise the new contact is ignored (for bucket updating purposes)
		} else {
			rt.AddContact(decodedData.Sender) //Adds contact to start of the bucket
		}



		if (decodedData.MessageType != "NONE" && decodedData.MessageType != "UNDEFINED"){
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
	
			responseRPC := NewRPC(rt.me, decodedData.Sender.Address, responseType, responseContent)
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

func (network *Network) SendPingMessage(message RPC) bool {
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

func (network *Network) SendFindContactMessage(contact *Contact) {
	// TODO
}

func (network *Network) SendFindDataMessage(hash string) {
	// TODO
}

func (network *Network) SendStoreMessage(data []byte) {
	// TODO
}
