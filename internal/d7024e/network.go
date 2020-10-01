package d7024e

import (
	"fmt"
	"net"
)

type Network struct {
	routingTable *RoutingTable
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

func Listen(ip string, port int) {
	PORT := ":8000" //Predefined port

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
		fmt.Print("RECEIVED: ", string(buffer[0:n-1]))

		//Update bucket corresponding to sender

		data := []byte("PONG")
		fmt.Printf("SENT: %s\n", string(data))
		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//func (network *Network) SendPingMessage(contact *Contact) {
func SendPingMessage(sender KademliaID, targetAddress string, targetPort string) {
	
	rpcMessage := NewRPC(sender, targetAddress, targetPort, "PING")
	fmt.Println("rpc created")
	fmt.Println(rpcMessage)
	//Convert rpcMessage to json or something before sending so it can be parsed properly
	
	CONNECT := targetAddress + ":" + targetPort

	s, err := net.ResolveUDPAddr("udp4", CONNECT)
	c, err := net.DialUDP("udp4", nil, s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer c.Close()

	for {
		data := []byte(rpcMessage)
		_, err = c.Write(data)

		if err != nil {
			fmt.Println(err)
			return
		}

		buffer := make([]byte, 1024)
		n, _, err := c.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("RECEIVED: %s\n", string(buffer[0:n]))
		//Update bucket appropriate to the recipient
		//Foreach bucket{Foreach contact{Look for IP adress}}
		//If the IP is found, put the contact at the end of the bucket
		//If it does not exist in a bucket, add it unless the bucket is full. (To which bucket?)
		//If the bucket is full, ping the contact at the top of the bucket. If that contact does not respond in a reasonable time it must be dropped and the new contact is added instead (but at the end of the list)
		//does bucket.AddContact() already do this??
		return
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
